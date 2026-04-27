package tui

import (
	"fmt"
	"strings"

	"github.com/awiipp/ranpo/internal/client"
	"github.com/awiipp/ranpo/internal/config"
	"github.com/awiipp/ranpo/internal/store"
	"github.com/awiipp/ranpo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
)

type CollectionModel struct {
	collections []string
	colCursor   int
	requests    []models.Request
	reqCursor   int
	pane        int // 0 = left (collections), 1 = right (requests)
	status      string
	err         error
}

func NewCollectionModel() CollectionModel {
	return CollectionModel{}
}

func (m CollectionModel) Init() tea.Cmd {
	return m.loadCollections()
}

func (m CollectionModel) Update(msg tea.Msg) (CollectionModel, tea.Cmd) {
	switch msg := msg.(type) {

	case collectionsLoadedMsg:
		m.collections = msg.names
		if len(m.collections) > 0 {
			return m, m.loadRequests(m.collections[0])
		}

	case requestsLoadedMsg:
		m.requests = msg.requests

	case responseMsg:
		if msg.err != nil {
			m.err = msg.err
		} else {
			return m, func() tea.Msg {
				return NavigateMsg{To: ScreenResponse, Response: msg.resp}
			}
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, func() tea.Msg { return NavigateMsg{To: ScreenHome} }

		case "tab":
			if len(m.collections) > 0 && len(m.requests) > 0 {
				m.pane = 1 - m.pane
			}

		case "up", "k":
			if m.pane == 0 {
				if m.colCursor > 0 {
					m.colCursor--
					return m, m.loadRequests(m.collections[m.colCursor])
				}
			} else {
				if m.reqCursor > 0 {
					m.reqCursor--
				}
			}

		case "down", "j":
			if m.pane == 0 {
				if m.colCursor < len(m.collections)-1 {
					m.colCursor++
					return m, m.loadRequests(m.collections[m.colCursor])
				}
			} else {
				if m.reqCursor < len(m.requests)-1 {
					m.reqCursor++
				}
			}

		case "enter", " ":
			if m.pane == 1 && len(m.requests) > 0 {
				return m, m.runRequest(m.requests[m.reqCursor])
			}

		case "d":
			if m.pane == 0 && len(m.collections) > 0 {
				name := m.collections[m.colCursor]
				_ = store.DeleteCollection(name)

				return m, m.loadCollections()
			}
		}
	}

	return m, nil
}

func (m CollectionModel) View() string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString("  " + titleStyle.Render("ranpo") + "  " + dimStyle.Render("collections") + "\n")
	sb.WriteString("  " + dimmerStyle.Render(strings.Repeat("─", 48)) + "\n\n")

	if len(m.collections) == 0 {
		sb.WriteString("  " + dimStyle.Render("no collections yet.") + "\n")
		sb.WriteString("  " + dimStyle.Render("use --save <name> when sending a request to create one.") + "\n")
	} else {
		// Two-pane layout
		leftWidth := 20
		colLines := m.renderCollections(leftWidth)
		reqLines := m.renderRequests()

		maxLen := max(len(colLines), len(reqLines))

		for i := range maxLen {
			left := ""
			if i < len(colLines) {
				left = colLines[i]
			}

			right := ""
			if i < len(reqLines) {
				right = reqLines[i]
			}

			leftPadded := fmt.Sprintf("%-*s", leftWidth+10, left)
			sb.WriteString("  " + leftPadded + "  " + right + "\n")
		}
	}

	if m.err != nil {
		sb.WriteString("\n  " + errorStyle.Render(m.err.Error()) + "\n")
	}

	sb.WriteString("\n")
	sb.WriteString(helpBar(
		"tab", "switch pane",
		"↑↓/jk", "navigate",
		"enter", "run request",
		"d", "delete collection",
		"esc", "back",
	))
	sb.WriteString("\n")

	return sb.String()
}

func (m CollectionModel) renderCollections(width int) []string {
	lines := []string{dimStyle.Render("collections")}

	for i, name := range m.collections {
		bullet := "  "
		style := normalItemStyle

		if i == m.colCursor {
			if m.pane == 0 {
				bullet = selectedItemStyle.Render("❯ ")
				style = selectedItemStyle
			} else {
				bullet = dimStyle.Render("› ")
				style = brightStyle
			}
		}
		lines = append(lines, bullet+style.Render(name))
	}

	return lines
}

func (m CollectionModel) renderRequests() []string {
	if len(m.requests) == 0 {
		return []string{dimStyle.Render("(empty)")}
	}

	lines := []string{dimStyle.Render("requests")}

	for i, req := range m.requests {
		bullet := "  "
		nameStyle := normalItemStyle

		if i == m.reqCursor {
			if m.pane == 1 {
				bullet = selectedItemStyle.Render("❯ ")
				nameStyle = selectedItemStyle
			} else {
				bullet = dimStyle.Render("› ")
				nameStyle = brightStyle
			}
		}

		line := bullet + methodBadge(req.Method) + " " + nameStyle.Render(req.Name)
		lines = append(lines, line)
	}

	return lines
}

type collectionsLoadedMsg struct{ names []string }
type requestsLoadedMsg struct{ requests []models.Request }

func (m CollectionModel) loadCollections() tea.Cmd {
	return func() tea.Msg {
		names, _ := store.ListCollections()
		return collectionsLoadedMsg{names: names}
	}
}

func (m CollectionModel) loadRequests(colName string) tea.Cmd {
	return func() tea.Msg {
		col, err := store.LoadCollection(colName)
		if err != nil {
			return requestsLoadedMsg{}
		}

		return requestsLoadedMsg{requests: col.Requests}
	}
}

func (m CollectionModel) runRequest(req models.Request) tea.Cmd {
	return func() tea.Msg {
		cfg, _ := config.Load()
		env, _ := store.LoadEnv(cfg.ActiveEnv)
		resp, err := client.Execute(&req, env)
		return responseMsg{resp: resp, err: err}
	}
}
