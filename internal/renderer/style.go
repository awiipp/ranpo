package renderer

import "github.com/charmbracelet/lipgloss"

var (
	keyStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("75"))  // blue
	strStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("114")) // green
	numStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("215")) // orange
	boolStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("204")) // pink
	nullStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("241")) // dim
	punctStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("250")) // light

	statusOKStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color("114")).Padding(0, 1)
	statusWarnStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color("215")).Padding(0, 1)
	statusErrStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color("204")).Padding(0, 1)
	dimStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	labelStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("75")).Bold(true)
)
