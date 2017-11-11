package client

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"time"

	gomem "github.com/bradfitz/gomemcache/memcache"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

type Codec struct {
	mc         *gomem.Client
	context    context.Context
	expiration time.Duration
	namespace  string
}

func NewCache(expire time.Duration, ns string) (*Codec, error) {
	var ctx context.Context
	var mc *gomem.Client
	switch appengine.IsDevAppServer() {
	case true:
		ctx = appengine.BackgroundContext()
	default:
		ctx = context.Background()
		mc = gomem.New("127.0.0.1:11211")
	}
	return &Codec{
		mc:         mc,
		context:    ctx,
		expiration: expire,
		namespace:  ns,
	}, nil
}

func (c *Codec) Get(key string, obj interface{}) error {
	var val []byte
	switch appengine.IsDevAppServer() {
	case true:
		item, err := memcache.Get(c.context, c.namespace+"_"+key)
		if err != nil {
			return err
		}
		val = item.Value
	default:
		item, err := c.mc.Get(c.namespace + "_" + key)
		if err != nil {
			return err
		}

		val = item.Value
	}

	err := c.Unmarshal(val, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *Codec) Add(key string, object interface{}) error {

	data, err := c.Marshal(object)
	if err != nil {
		return err
	}

	switch appengine.IsDevAppServer() {
	case true:
		item := memcache.Item{
			Expiration: c.expiration,
			Key:        c.namespace + "_" + key,
			Object:     object,
			Value:      data,
		}

		return memcache.Add(c.context, &item)
	default:
		item := gomem.Item{
			Expiration: int32(c.expiration.Seconds()),
			Key:        c.namespace + "_" + key,
			Value:      data,
		}
		return c.mc.Add(&item)
	}
}

func (c *Codec) Marshal(obj interface{}) ([]byte, error) {
	network := new(bytes.Buffer)
	enc := gzip.NewWriter(network)

	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	_, err = enc.Write(data)
	if err != nil {
		return nil, err
	}

	return network.Bytes(), nil
}

func (c *Codec) Unmarshal(data []byte, obj interface{}) error {
	network := new(bytes.Buffer)
	dec, err := gzip.NewReader(network)
	switch {
	case err == io.EOF:
		return nil
	case err != nil:
		return err
	}

	_, err = dec.Read(data)
	if err != nil {
		return err
	}

	return json.NewDecoder(network).Decode(&obj)
}
