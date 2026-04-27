package tui

import "github.com/awiipp/ranpo/pkg/models"

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
