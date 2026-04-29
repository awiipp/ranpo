package tui

import tea "github.com/charmbracelet/bubbletea"

func Launch() error {
	p := tea.NewProgram(
		NewApp(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	_, err := p.Run()
	return err
}

func LaunchWithRequest(method, url string) error {
	app := NewApp()
	app.screen = ScreenRequestForm
	form := NewRequestFormModel(method)
	form.urlInput.SetValue(url)
	form.urlInput.CursorEnd()
	app.requestForm = form

	p := tea.NewProgram(
		app,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	_, err := p.Run()
	return err
}
