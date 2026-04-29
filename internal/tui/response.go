package tui

import (
	"fmt"
	"strings"

	"github.com/awiipp/ranpo/internal/renderer"
	"github.com/awiipp/ranpo/pkg/models"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type ResponseModel struct {
	resp     *models.Response
	viewport viewport.Model
	ready    bool
	tab      int // 0 = body, 1 = headers
}

func NewResponseModel(resp *models.Response, w int, h int) ResponseModel {
	vp := viewport.New(w-4, h-10)
	m := ResponseModel{resp: resp, viewport: vp, ready: true}

	m.viewport.SetContent(m.buildContent())
	return m
}

func (m ResponseModel) Init() tea.Cmd { return nil }

func (m ResponseModel) Update(msg tea.Msg) (ResponseModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width - 4
		m.viewport.Height = msg.Height - 6
		m.viewport.SetContent(m.buildContent())

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, navCmd(ScreenHome, "")
		case "b":
			return m, navCmd(ScreenRequestForm, "")
		case "1":
			m.tab = 0
			m.viewport.SetContent(m.buildContent())
		case "2":
			m.tab = 1
			m.viewport.SetContent(m.buildContent())
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m ResponseModel) buildContent() string {
	if m.resp == nil {
		return dimStyle.Render("no response")
	}

	if m.tab == 1 {
		return renderer.HeadersView(m.resp.Headers)
	}

	return renderer.PrettyJSON(m.resp.Body)
}

func (m ResponseModel) View() string {
	if m.resp == nil {
		return "\n  " + dimStyle.Render("no response to display") + "\n"
	}

	var sb strings.Builder
	sb.WriteString("\n")

	// Status line
	sb.WriteString("  ")
	sb.WriteString(statusBadge(m.resp.StatusCode))
	sb.WriteString("  " + dimStyle.Render(m.resp.Status))
	sb.WriteString("  " + dimStyle.Render(m.resp.Duration.String()))
	sb.WriteString("  " + dimStyle.Render(fmt.Sprintf("%d bytes", len(m.resp.Body))))
	sb.WriteString("\n")

	// Tab bar
	b := mutedStyle.Render("body")
	h := mutedStyle.Render("headers")

	if m.tab == 0 {
		b = labelStyle.Render("body")
	} else {
		h = labelStyle.Render("headers")
	}

	sb.WriteString("         view:  " + b + dimStyle.Render(" [1]") + "  " + h + dimStyle.Render(" [2]") + "\n")
	sb.WriteString("  " + dividerLine(70) + "\n\n")

	const bodyIndent = "  "
	sb.WriteString(renderer.IndentBlock(m.viewport.View(), bodyIndent) + "\n\n")

	pct := int(m.viewport.ScrollPercent() * 100)

	sb.WriteString(helpBar("↑↓ PgUp PgDn", "scroll", "1/2", "body/headers", "b", "back to form", "esc", "home"))
	sb.WriteString("  " + dimmerStyle.Render(fmt.Sprintf("%d%%", pct)) + "\n")

	return sb.String()
}
