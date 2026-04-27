package main

import "github.com/awiipp/ranpo/cmd"

func main() {
	cmd.Execute()

	// p := tea.NewProgram(tui.NewApp())

	// if _, err := p.Run(); err != nil {
	// 	log.Fatal(err)
	// }

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
