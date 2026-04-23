package models

// Environment holds a named set of key-value variables.
type Environtment struct {
	Name      string            `json:"name"`
	Variables map[string]string `json:"variables"`
}
