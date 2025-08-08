package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response wraps http.Response to provide convenient methods for handling responses
type Response struct {
	*http.Response
	body []byte
}

// newResponse creates a new Response instance
func newResponse(resp *http.Response) (*Response, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return &Response{
		Response: resp,
		body:     body,
	}, nil
}

// StatusCode returns the HTTP status code
func (r *Response) StatusCode() int {
	return r.Response.StatusCode
}

// IsSuccess returns true if the status code is in the 2xx range
func (r *Response) IsSuccess() bool {
	return r.StatusCode() >= 200 && r.StatusCode() < 300
}

// IsError returns true if the status code is in the 4xx or 5xx range
func (r *Response) IsError() bool {
	return r.StatusCode() >= 400
}

// IsClientError returns true if the status code is in the 4xx range
func (r *Response) IsClientError() bool {
	return r.StatusCode() >= 400 && r.StatusCode() < 500
}

// IsServerError returns true if the status code is in the 5xx range
func (r *Response) IsServerError() bool {
	return r.StatusCode() >= 500
}

// String returns the response body as a string
func (r *Response) String() string {
	return string(r.body)
}

// Bytes returns the response body as bytes
func (r *Response) Bytes() []byte {
	return r.body
}

// Decode decodes the response body into the given interface
func (r *Response) Decode(v any) error {
	if len(r.body) == 0 {
		return fmt.Errorf("response body is empty")
	}

	return json.Unmarshal(r.body, v)
}

// DecodeJSON is an alias for Decode for better readability
func (r *Response) DecodeJSON(v any) error {
	return r.Decode(v)
}

// GetHeader returns the first value of the specified header
func (r *Response) GetHeader(key string) string {
	return r.Header.Get(key)
}

// GetHeaders returns all values of the specified header
func (r *Response) GetHeaders(key string) []string {
	return r.Header.Values(key)
}

// ContentType returns the Content-Type header
func (r *Response) ContentType() string {
	return r.GetHeader("Content-Type")
}

// Location returns the Location header (useful for redirects)
func (r *Response) Location() string {
	return r.GetHeader("Location")
}

// Error returns an error if the response indicates an error
func (r *Response) Error() error {
	if r.IsError() {
		return fmt.Errorf("HTTP %d: %s", r.StatusCode(), r.String())
	}
	return nil
}

// MustString returns the response body as a string, panicking if there's an error
func (r *Response) MustString() string {
	return r.String()
}

// MustDecode decodes the response body into the given interface, panicking if there's an error
func (r *Response) MustDecode(v any) {
	if err := r.Decode(v); err != nil {
		panic(fmt.Sprintf("failed to decode response: %v", err))
	}
}
