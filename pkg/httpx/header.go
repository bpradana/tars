package httpx

import (
	"net/http"
)

// Header wraps http.Header to provide convenient methods for common headers
type Header struct {
	headers http.Header
}

// NewHeader creates a new Header instance
func NewHeader() *Header {
	return &Header{
		headers: make(http.Header),
	}
}

// Add adds a header with the given key and value
func (h *Header) Add(key, value string) *Header {
	h.headers.Add(key, value)
	return h
}

// Set sets a header with the given key and value (replaces existing values)
func (h *Header) Set(key, value string) *Header {
	h.headers.Set(key, value)
	return h
}

// Authorization adds an Authorization header
func (h *Header) Authorization(token string) *Header {
	return h.Set("Authorization", token)
}

// Bearer adds a Bearer token Authorization header
func (h *Header) Bearer(token string) *Header {
	return h.Set("Authorization", "Bearer "+token)
}

// ContentType sets the Content-Type header
func (h *Header) ContentType(contentType string) *Header {
	return h.Set("Content-Type", contentType)
}

// JSON sets the Content-Type header to application/json
func (h *Header) JSON() *Header {
	return h.ContentType("application/json")
}

// XML sets the Content-Type header to application/xml
func (h *Header) XML() *Header {
	return h.ContentType("application/xml")
}

// Form sets the Content-Type header to application/x-www-form-urlencoded
func (h *Header) Form() *Header {
	return h.ContentType("application/x-www-form-urlencoded")
}

// UserAgent sets the User-Agent header
func (h *Header) UserAgent(agent string) *Header {
	return h.Set("User-Agent", agent)
}

// Accept sets the Accept header
func (h *Header) Accept(accept string) *Header {
	return h.Set("Accept", accept)
}

// AcceptJSON sets the Accept header to application/json
func (h *Header) AcceptJSON() *Header {
	return h.Accept("application/json")
}

// AcceptXML sets the Accept header to application/xml
func (h *Header) AcceptXML() *Header {
	return h.Accept("application/xml")
}

// Get returns the header values for the given key
func (h *Header) Get(key string) []string {
	return h.headers.Values(key)
}

// GetFirst returns the first header value for the given key
func (h *Header) GetFirst(key string) string {
	return h.headers.Get(key)
}

// Delete removes the header with the given key
func (h *Header) Delete(key string) *Header {
	h.headers.Del(key)
	return h
}

// Headers returns the underlying http.Header
func (h *Header) Headers() http.Header {
	return h.headers
}

// Clone creates a copy of the header
func (h *Header) Clone() *Header {
	clone := NewHeader()
	for key, values := range h.headers {
		for _, value := range values {
			clone.Add(key, value)
		}
	}
	return clone
}
