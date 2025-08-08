package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Request wraps http.Request to provide convenient methods for building requests
type Request struct {
	*http.Request
	client *http.Client
}

// newRequest creates a new Request instance
func newRequest(method, url string, client *http.Client) (*Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return &Request{
		Request: req,
		client:  client,
	}, nil
}

// WithHeader adds a header to the request
func (r *Request) WithHeader(key, value string) *Request {
	r.Header.Set(key, value)
	return r
}

// WithHeaders adds multiple headers to the request
func (r *Request) WithHeaders(headers map[string]string) *Request {
	for key, value := range headers {
		r.Header.Set(key, value)
	}
	return r
}

// WithHeaderObject adds headers from a Header object
func (r *Request) WithHeaderObject(h *Header) *Request {
	for key, values := range h.headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
	return r
}

// WithBody sets the request body
func (r *Request) WithBody(body io.Reader) *Request {
	r.Body = io.NopCloser(body)
	return r
}

// WithJSON sets the request body to JSON and sets Content-Type header
func (r *Request) WithJSON(data any) *Request {
	jsonData, err := json.Marshal(data)
	if err != nil {
		// In a real implementation, you might want to handle this error differently
		panic(fmt.Sprintf("failed to marshal JSON: %v", err))
	}

	r.Header.Set("Content-Type", "application/json")
	r.Body = io.NopCloser(bytes.NewBuffer(jsonData))
	return r
}

// WithForm sets the request body to form data and sets Content-Type header
func (r *Request) WithForm(data map[string]string) *Request {
	values := url.Values{}
	for key, value := range data {
		values.Set(key, value)
	}

	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Body = io.NopCloser(strings.NewReader(values.Encode()))
	return r
}

// WithQuery adds query parameters to the request URL
func (r *Request) WithQuery(params map[string]string) *Request {
	q := r.URL.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	r.URL.RawQuery = q.Encode()
	return r
}

// WithQueryParam adds a single query parameter to the request URL
func (r *Request) WithQueryParam(key, value string) *Request {
	q := r.URL.Query()
	q.Set(key, value)
	r.URL.RawQuery = q.Encode()
	return r
}

// WithTimeout sets a timeout for the request (requires a custom client)
func (r *Request) WithTimeout(timeout time.Duration) *Request {
	if r.client == nil {
		r.client = &http.Client{}
	}
	r.client.Timeout = timeout
	return r
}

// Do executes the request and returns a Response
func (r *Request) Do() (*Response, error) {
	client := r.client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(r.Request)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return newResponse(resp)
}

// MustDo executes the request and returns a Response, panicking if there's an error
func (r *Request) MustDo() *Response {
	resp, err := r.Do()
	if err != nil {
		panic(fmt.Sprintf("failed to execute request: %v", err))
	}
	return resp
}
