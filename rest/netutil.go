package rest

import (
	"net/http"
	"net/url"

	"github.com/akshith-gunasheelan/terraform-provider-com/utils/"
)

type Options struct {
	Headers map[string]string
	Query   map[string]interface{}
}
type Client struct {
	ClientId     string
	ClientSecret string
	Endpoint     string
	ApiVersion   string
	Option       Options
}

func (c *Client) NewClient(clientid, clientsecret, endpoint string) *Client {
	return &Client{ClientId: clientid, ClientSecret: clientsecret, Endpoint: endpoint, Option: Options{}}
}
func (c *Client) SetAuthHeaderOptions(headers map[string]string) {
	c.Option.Headers = headers
}
func (c *Client) GetQueryStrings(u *url.URL, query map[string]interface{}) {
	if len(query) == 0 {
		return
	}
	parameters := url.Values{}
	for k, v := range query {
		if val, ok := v.([]string); ok {
			for _, va := range val {
				parameters.Add(k, va)
			}
		} else {
			parameters.Add(k, v.(string))
		}
		u.RawQuery = parameters.Encode()
	}
	return
}

// GetQueryString - get a query string for url through the Client Struct
func (c *Client) GetQueryString(u *url.URL) {
	if len(c.Option.Query) == 0 {
		return
	}
	parameters := url.Values{}
	for k, v := range c.Option.Query {
		if val, ok := v.([]string); ok {
			for _, va := range val {
				parameters.Add(k, va)
			}
		} else {
			parameters.Add(k, v.(string))
		}
		u.RawQuery = parameters.Encode()
	}
	return
}
func (c *Client) RestAPICall(method Method, path string, options interface{}, query ...map[string]interface{}) ([]byte, error) {
	var (
		Url *url.URL
		err error
		req *http.Request
	)
	Url, err = url.Parse(utils.Sanatize(c.Endpoint))
	if err != nil {
		return nil, err
	}
	Url.Path += path
}

