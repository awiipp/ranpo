package models

// Collection is a named group of saved requests.
type Collection struct {
	Name     string    `json:"name"`
	Requests []Request `json:"requests"`
}
