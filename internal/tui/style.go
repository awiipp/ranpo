package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	colorPurple    = "135"
	colorBlue      = "75"
	colorGreen     = "114"
	colorOrange    = "215"
	colorRed       = "204"
	colorYellow    = "221"
	colorDim       = "241"
	colorDimmer    = "238"
	colorMuted     = "245"
	colorBright    = "255"
	colorBorder    = "237"
	colorBorderFoc = "75"
)

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colorPurple))
	dimStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(colorDim))
	dimmerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colorDimmer))
	mutedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color(colorMuted))
	brightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colorBright))
	labelStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colorBlue))

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(colorBorderFoc)).
				Padding(0, 1)

	blurredBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(colorBorder)).
				Padding(0, 1)

	statusOKStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colorGreen))
	statusWarnStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colorOrange))
	statusErrorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colorRed))
	successStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color(colorGreen))
	errorStyle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colorRed))

	selectedItemStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(colorBlue))
	normalItemStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color(colorMuted))
	helpStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color(colorDimmer))

	methodColors = map[string]string{
		"GET":    colorGreen,
		"POST":   colorBlue,
		"PUT":    colorYellow,
		"PATCH":  colorOrange,
		"DELETE": colorRed,
	}
)

func methodBadge(method string) string {
	col, ok := methodColors[method]
	if !ok {
		return mutedStyle.Render("[" + method + "]")
	}

	return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(col)).Render("[" + method + "]")
}

func statusBadge(code int) string {
	s := fmt.Sprintf("%d", code)

	switch {
	case code >= 200 && code < 300:
		return statusOKStyle.Render(s)
	case code >= 300 && code < 400:
		return statusWarnStyle.Render(s)
	default:
		return statusErrorStyle.Render(s)
	}
}

func helpBar(pairs ...string) string {
	var parts []string
	for i := 0; i+1 < len(pairs); i += 2 {
		parts = append(parts, dimStyle.Render(": ")+mutedStyle.Render(pairs[i+1]))
	}

	return helpStyle.Render("  "+strings.Join(parts, "   ")) + "\n"
}

func dividerLine(width int) string {
	if width <= 0 {
		width = 48
	}
	return dimmerStyle.Render(strings.Repeat("─", width))
}
