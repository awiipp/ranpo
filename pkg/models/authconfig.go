package models

// AuthConfig holds authentication configuration for a request.
type AuthConfig struct {
	Type  string `json:"type"`
	Token string `json:"token,omitempty"`
	User  string `json:"user,omitempty"`
	Pass  string `json:"pass,omitempty"`
}
