package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Provider interface {
	Get(Options) error
	Post(Options) error
	Put(Options) error
	Patch(Options) error
	Head(Options) error
	Delete(Options) error
	Options(Options) error
}

type Client struct {
	Hostname   string
	Schema     string
	PublicKey  string
	PrivateKey string
	Client     *http.Client
}

type Error struct {
	Message string `json:"message"`
	Details string `json:"messageDetails"`
}

func (c *Client) Get(opts Options) error {
	if c.Client == nil {
		return errors.New("invalid client.Client")
	}

	if opts.QueryString == nil {
		opts.QueryString = make(map[string]string, 0)
	}
	opts.QueryString["key"] = c.PublicKey

	path, err := c.url(opts)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	resp, err := c.Client.Do(req)
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

func (c *Client) Post(opts Options) error {
	panic("not implemented")
}

func (c *Client) Put(opts Options) error {
	panic("not implemented")
}

func (c *Client) Patch(opts Options) error {
	panic("not implemented")
}

func (c *Client) Head(opts Options) error {
	panic("not implemented")
}

func (c *Client) Delete(opts Options) error {
	panic("not implemented")
}

func (c *Client) Options(opts Options) error {
	panic("not implemented")
}

func (c Client) url(o Options) (string, error) {
	if c.Schema == "" {
		return "", errors.New("must specify a request schema 'http|https'")
	}
	if c.Hostname == "" {
		return "", errors.New("must specify a host to make requests against")
	}

	if strings.HasPrefix(o.Endpoint, "/") {
		o.Endpoint = strings.TrimPrefix(o.Endpoint, "/")
	}

	path := fmt.Sprintf(
		"%s://%s/%s",
		c.Schema,
		c.Hostname,
		o.Endpoint,
	)

	if o.QueryString != nil {
		var qs string
		for k, v := range o.QueryString {
			qs = fmt.Sprintf("%s%s=%s&", qs, k, v)
		}
		path = fmt.Sprintf("%s?%s", path, qs)
	}

	return path, nil
}
