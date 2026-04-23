package models

// Collection is a named group of saved requests.
type Collection struct {
	Name    string    `json:"name"`
	Request []Request `json:"request"`
}
