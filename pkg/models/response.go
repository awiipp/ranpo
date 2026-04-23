package models

import "time"

// Response is the result of an executed HTTP request.
type Response struct {
	StatusCode int
	Status     string
	Headers    map[string][]string
	Body       []byte
	Duration   time.Duration
}
