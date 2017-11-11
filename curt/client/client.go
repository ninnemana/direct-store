package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/logging"
	gomem "github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/urlfetch"
)

type Provider interface {
	Do(Options) error
	// Get(Options) error
	// Post(Options) error
	// Put(Options) error
	// Patch(Options) error
	// Head(Options) error
	// Delete(Options) error
	// Options(Options) error
}

type Client struct {
	hostname   string
	schema     string
	publicKey  string
	privateKey string
	client     *http.Client
	cache      *Codec
	log        *logging.Logger
}

type Error struct {
	Message string `json:"message"`
	Details string `json:"messageDetails"`
}

func New(l *logging.Logger, schema, host, pubKey, privateKey string) (*Client, error) {
	if l == nil {
		return nil, errors.New("missing logging.Logger")
	}

	cache, err := NewCache(time.Hour*24, "directstore")
	if err != nil {
		return nil, err
	}

	client := Client{
		log:        l,
		hostname:   host,
		schema:     schema,
		publicKey:  pubKey,
		privateKey: privateKey,
		cache:      cache,
	}

	switch appengine.IsDevAppServer() {
	case false:
		client.client = http.DefaultClient
	default:
		client.client = urlfetch.Client(context.Background())
	}

	l.Log(logging.Entry{
		Severity: logging.Debug,
		Payload:  "Creating new CURT HTTP Client",
		Labels: map[string]string{
			"pkg":       "curt/client",
			"func":      "new",
			"hostname":  host,
			"schema":    schema,
			"publicKey": pubKey,
		},
	})

	return &client, nil
}

func (c *Client) RequestKey(opts Options) (string, error) {
	path, err := c.url(opts)
	if err != nil {
		return "", err
	}
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	qs, err := opts.QueryString.Build()
	if err != nil {
		return "", err
	}

	endpoint := strings.Replace(opts.Endpoint, "/", "_", -1)
	query := []string{}
	for k, v := range qs {
		query = append(query, fmt.Sprintf("%s:%s", k, v))
	}

	key := fmt.Sprintf(
		"%s_%s",
		u.Host,
		endpoint,
	)

	if len(query) > 0 {
		key = fmt.Sprintf(
			"%s_%s",
			key,
			strings.Join(query, "_"),
		)
	}

	return key, nil
}

func (c *Client) Do(opts Options) error {
	cacheKey, err := c.RequestKey(opts)
	if err != nil {
		return errors.Wrap(err, "failed to create cache key")
	}

	var item interface{}
	err = c.cache.Get(cacheKey, &item)
	switch err {
	case nil:
		opts.Result = item
		return nil
	case memcache.ErrCacheMiss:
	case gomem.ErrCacheMiss:
	default:
		// c.log.Log(logging.Entry{
		// 	Severity: logging.Error,
		// 	Payload: map[string]interface{}{
		// 		"msg": "Failed to make memcache GET call",
		// 		"err": err.Error(),
		// 		"key": cacheKey,
		// 	},
		// 	Labels: map[string]string{
		// 		"pkg":  "curt/client",
		// 		"func": "Do",
		// 	},
		// })
		// return errors.Wrap(err, "failed to make memcache call")
	}

	if c.client == nil {
		return errors.New("invalid client.Client")
	}

	switch opts.Method {
	case http.MethodGet:
		err = c.get(opts)
	case http.MethodPost:
		err = c.post(opts)
	case http.MethodPut:
		err = c.put(opts)
	case http.MethodPatch:
		err = c.patch(opts)
	case http.MethodDelete:
		err = c.delete(opts)
	// case http.MethodHead:
	// case http.MethodOptions:
	// case http.MethodTrace:
	// case http.MethodConnect:
	default:
		return errors.Errorf("HTTP Method (%s) not supported", opts.Method)
	}

	if err != nil {
		return errors.Wrap(err, "failed to make client HTTP call")
	}

	return nil
	// return c.cache.Add(cacheKey, &opts.Result)
}

func (c *Client) get(opts Options) error {

	path, err := c.url(opts)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decdr := json.NewDecoder(resp.Body)

	switch {
	case resp.StatusCode < 400:
		err := decdr.Decode(&opts.Result)
		if err != nil {
			return err
		}
	default:
		var e Error
		err := decdr.Decode(&e)
		if err != nil {
			return err
		}

		err = errors.New(e.Message)
		return errors.Wrapf(err, "failed to make HTTP request '%s'", e.Details)
	}

	return nil
}

func (c *Client) post(opts Options) error {
	panic("not implemented")
}

func (c *Client) put(opts Options) error {
	panic("not implemented")
}

func (c *Client) patch(opts Options) error {
	panic("not implemented")
}

func (c *Client) head(opts Options) error {
	panic("not implemented")
}

func (c *Client) delete(opts Options) error {
	panic("not implemented")
}

func (c *Client) options(opts Options) error {
	panic("not implemented")
}

func (c Client) url(o Options) (string, error) {
	if c.schema == "" {
		return "", errors.New("must specify a request schema 'http|https'")
	}
	if c.hostname == "" {
		return "", errors.New("must specify a host to make requests against")
	}

	if strings.HasPrefix(o.Endpoint, "/") {
		o.Endpoint = strings.TrimPrefix(o.Endpoint, "/")
	}

	path := fmt.Sprintf(
		"%s://%s/%s",
		c.schema,
		c.hostname,
		o.Endpoint,
	)

	qs := map[string]string{}
	if o.QueryString != nil {
		var err error
		qs, err = o.QueryString.Build()
		if err != nil {
			return "", err
		}
	}

	qs["key"] = c.publicKey

	if qs != nil {
		var query string
		for k, v := range qs {
			query = fmt.Sprintf("%s%s=%s&", query, k, v)
		}
		path = fmt.Sprintf("%s?%s", path, query)
	}

	return path, nil
}
