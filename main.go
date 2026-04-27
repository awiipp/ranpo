package main

import (
	"log"

	"github.com/awiipp/ranpo/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.NewApp())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	// body := []byte(`{
	// 	"name": "Nelsi Cornelia",
	// 	"age": 18,
	// 	"cute": true
	// 	}`)

	// output := renderer.RenderResponse(
	// 	200,
	// 	"200 OK",
	// 	body,
	// 	time.Since(time.Now().Add(-150*time.Millisecond)),
	// )

	// fmt.Println(output)
}
