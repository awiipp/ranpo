package models

// Request represents a single HTTP request, saved or in-flight.
type Request struct {
	Name    string            `json:"name"`
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
	Auth    *AuthConfig       `json:"auth,omitempty"` // pointer so that can use omitempty.
}
