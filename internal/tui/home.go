package tui

import tea "github.com/charmbracelet/bubbletea"

type menuItem struct {
	key    string
	label  string
	desc   string
	screen Screen
	method string
}

var menuItems = []menuItem{
	{"g", "GET", "send a GET request", ScreenRequestForm, "GET"},
	{"p", "POST", "send a POST request with body", ScreenRequestForm, "POST"},
	{"u", "PUT", "update a resource", ScreenRequestForm, "PUT"},
	{"P", "PATCH", "partial update", ScreenRequestForm, "PATCH"},
	{"d", "DELETE", "delete a resource", ScreenRequestForm, "DELETE"},
	{"c", "Collections", "browse & run saved requests", ScreenCollections, ""},
	{"e", "Environments", "manage dev / staging / prod", ScreenEnvManager, ""},
}

type HomeModel struct {
	cursor int
}

func NewHomeModel() HomeModel {
	return HomeModel{}
}

func (m HomeModel) Init() tea.Cmd {
	return nil
}

func (m HomeModel) Update(msg tea.Msg) (HomeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(menuItems)-1 {
				m.cursor++
			}

		case "enter", " ":
			item := menuItems[m.cursor]
			return m, navCmd(item.screen, item.method)

		case "q":
			return m, tea.Quit

		default:
			for _, item := range menuItems {
				if msg.String() == item.key {
					return m, navCmd(item.screen, item.method)
				}
			}
		}
	}

	return m, nil
}

func navCmd(screen Screen, method string) tea.Cmd {
	return func() tea.Msg {
		return NavigateMsg{To: screen, Method: method}
	}
}
