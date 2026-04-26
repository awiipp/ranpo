package main

import (
	"fmt"
	"time"

	"github.com/awiipp/ranpo/internal/renderer"
)

func main() {
	body := []byte(`{
		"name": "Nelsi Cornelia",
		"age": 18,
		"cute": true
		}`)

	output := renderer.RenderResponse(
		200,
		"200 OK",
		body,
		time.Since(time.Now().Add(-150*time.Millisecond)),
	)

	fmt.Println(output)
}
