package tui

import (
	"github.com/awiipp/ranpo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen int

const (
	ScreenHome Screen = iota
	ScreenRequestForm
	ScreenResponse
	ScreenCollections
	ScreenEnvManager
)

type NavigateMsg struct {
	To       Screen
	Method   string
	Response *models.Response
}

type AppModel struct {
	screen      Screen
	width       int
	height      int
	home        HomeModel
	requestForm RequestFormModel
	response    ResponseModel
	collections CollectionModel
	envManager  EnvModel
}

func NewApp() AppModel {
	return AppModel{
		screen:      ScreenHome,
		home:        NewHomeModel(),
		requestForm: NewRequestFormModel("GET"),
		collections: NewCollectionModel(),
		envManager:  NewEnvModel(),
	}
}

func (m AppModel) Init() tea.Cmd { return nil }

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		if m.response.ready {
			m.response.viewport.Width = msg.Width - 4
			m.response.viewport.Height = msg.Height - 6
		}

	case NavigateMsg:
		return m.handleNavigate(msg)

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	switch m.screen {
	case ScreenHome:
		m.home, cmd = m.home.Update(msg)
	case ScreenRequestForm:
		m.requestForm, cmd = m.requestForm.Update(msg)
	case ScreenResponse:
		m.response, cmd = m.response.Update(msg)
	case ScreenCollections:
		m.collections, cmd = m.collections.Update(msg)
	case ScreenEnvManager:
		m.envManager, cmd = m.envManager.Update(msg)
	}

	return m, cmd
}

func (m AppModel) handleNavigate(msg NavigateMsg) (tea.Model, tea.Cmd) {
	m.screen = msg.To
	switch msg.To {
	case ScreenHome:
		m.home = NewHomeModel()
		return m, nil

	case ScreenRequestForm:
		m.requestForm = NewRequestFormModel(msg.Method)
		return m, m.requestForm.Init()

	case ScreenResponse:
		m.response = NewResponseModel(msg.Response, m.width, m.height)
		return m, nil

	case ScreenCollections:
		m.collections = NewCollectionModel()
		return m, m.collections.Init()

	case ScreenEnvManager:
		m.envManager = NewEnvModel()
		return m, m.envManager.Init()
	}

	return m, nil
}

func (m AppModel) View() string {
	switch m.screen {
	case ScreenHome:
		return m.home.View()
	case ScreenRequestForm:
		return m.requestForm.View()
	case ScreenResponse:
		return m.response.View()
	case ScreenCollections:
		return m.collections.View()
	case ScreenEnvManager:
		return m.envManager.View()
	}
	return ""
}

func Launch() error {
	p := tea.NewProgram(
		NewApp(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	_, err := p.Run()
	return err
}
