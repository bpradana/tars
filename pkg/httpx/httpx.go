package httpx

import (
	"net/http"
	"time"
)

// Client represents an HTTP client with default settings
type Client struct {
	httpClient     *http.Client
	baseURL        string
	defaultHeaders *Header
}

// NewClient creates a new HTTP client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		defaultHeaders: NewHeader(),
	}
}

// WithBaseURL sets the base URL for all requests
func (c *Client) WithBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

// WithTimeout sets the default timeout for requests
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	c.httpClient.Timeout = timeout
	return c
}

// WithDefaultHeaders sets default headers for all requests
func (c *Client) WithDefaultHeaders(headers *Header) *Client {
	c.defaultHeaders = headers
	return c
}

// WithDefaultHeader adds a default header for all requests
func (c *Client) WithDefaultHeader(key, value string) *Client {
	c.defaultHeaders.Set(key, value)
	return c
}

// buildURL constructs the full URL for a request
func (c *Client) buildURL(path string) string {
	if c.baseURL != "" {
		return c.baseURL + path
	}
	return path
}

// createRequest creates a new request with default headers
func (c *Client) createRequest(method, url string) (*Request, error) {
	req, err := newRequest(method, url, c.httpClient)
	if err != nil {
		return nil, err
	}

	// Add default headers
	if c.defaultHeaders != nil {
		req.WithHeaderObject(c.defaultHeaders)
	}

	return req, nil
}

// GET performs a GET request
func (c *Client) GET(url string) (*Request, error) {
	return c.createRequest(http.MethodGet, c.buildURL(url))
}

// POST performs a POST request
func (c *Client) POST(url string) (*Request, error) {
	return c.createRequest(http.MethodPost, c.buildURL(url))
}

// PUT performs a PUT request
func (c *Client) PUT(url string) (*Request, error) {
	return c.createRequest(http.MethodPut, c.buildURL(url))
}

// DELETE performs a DELETE request
func (c *Client) DELETE(url string) (*Request, error) {
	return c.createRequest(http.MethodDelete, c.buildURL(url))
}

// PATCH performs a PATCH request
func (c *Client) PATCH(url string) (*Request, error) {
	return c.createRequest(http.MethodPatch, c.buildURL(url))
}

// HEAD performs a HEAD request
func (c *Client) HEAD(url string) (*Request, error) {
	return c.createRequest(http.MethodHead, c.buildURL(url))
}

// OPTIONS performs an OPTIONS request
func (c *Client) OPTIONS(url string) (*Request, error) {
	return c.createRequest(http.MethodOptions, c.buildURL(url))
}

// Convenience functions for common HTTP operations

// Get performs a GET request and returns the response
func (c *Client) Get(url string) (*Response, error) {
	req, err := c.GET(url)
	if err != nil {
		return nil, err
	}
	return req.Do()
}

// Post performs a POST request with JSON body and returns the response
func (c *Client) Post(url string, data any) (*Response, error) {
	req, err := c.POST(url)
	if err != nil {
		return nil, err
	}
	return req.WithJSON(data).Do()
}

// PostForm performs a POST request with form data and returns the response
func (c *Client) PostForm(url string, data map[string]string) (*Response, error) {
	req, err := c.POST(url)
	if err != nil {
		return nil, err
	}
	return req.WithForm(data).Do()
}

// Put performs a PUT request with JSON body and returns the response
func (c *Client) Put(url string, data any) (*Response, error) {
	req, err := c.PUT(url)
	if err != nil {
		return nil, err
	}
	return req.WithJSON(data).Do()
}

// Delete performs a DELETE request and returns the response
func (c *Client) Delete(url string) (*Response, error) {
	req, err := c.DELETE(url)
	if err != nil {
		return nil, err
	}
	return req.Do()
}

// Patch performs a PATCH request with JSON body and returns the response
func (c *Client) Patch(url string, data any) (*Response, error) {
	req, err := c.PATCH(url)
	if err != nil {
		return nil, err
	}
	return req.WithJSON(data).Do()
}

// Global client instance for convenience
var defaultClient = NewClient()

// Global convenience functions

// GET performs a GET request using the default client
func GET(url string) (*Request, error) {
	return defaultClient.GET(url)
}

// POST performs a POST request using the default client
func POST(url string) (*Request, error) {
	return defaultClient.POST(url)
}

// PUT performs a PUT request using the default client
func PUT(url string) (*Request, error) {
	return defaultClient.PUT(url)
}

// DELETE performs a DELETE request using the default client
func DELETE(url string) (*Request, error) {
	return defaultClient.DELETE(url)
}

// PATCH performs a PATCH request using the default client
func PATCH(url string) (*Request, error) {
	return defaultClient.PATCH(url)
}

// Get performs a GET request using the default client and returns the response
func Get(url string) (*Response, error) {
	return defaultClient.Get(url)
}

// Post performs a POST request with JSON body using the default client and returns the response
func Post(url string, data any) (*Response, error) {
	return defaultClient.Post(url, data)
}

// PostForm performs a POST request with form data using the default client and returns the response
func PostForm(url string, data map[string]string) (*Response, error) {
	return defaultClient.PostForm(url, data)
}

// Put performs a PUT request with JSON body using the default client and returns the response
func Put(url string, data any) (*Response, error) {
	return defaultClient.Put(url, data)
}

// Delete performs a DELETE request using the default client and returns the response
func Delete(url string) (*Response, error) {
	return defaultClient.Delete(url)
}

// Patch performs a PATCH request with JSON body using the default client and returns the response
func Patch(url string, data any) (*Response, error) {
	return defaultClient.Patch(url, data)
}
