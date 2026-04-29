package tui

import (
	"fmt"
	"strings"

	"github.com/awiipp/ranpo/internal/client"
	"github.com/awiipp/ranpo/internal/config"
	"github.com/awiipp/ranpo/internal/store"
	"github.com/awiipp/ranpo/pkg/models"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type formField int

const (
	fieldURL formField = iota
	fieldToken
	fieldBody
	fieldSave
	fieldCount
)

type responseMsg struct {
	resp *models.Response
	err  error
}

type RequestFormModel struct {
	method     string
	urlInput   textinput.Model
	tokenInput textinput.Model
	saveInput  textinput.Model
	bodyArea   textarea.Model
	focused    formField
	loading    bool
	savedName  string
	err        error
}

func NewRequestFormModel(method string) RequestFormModel {
	url := textinput.New()
	url.Placeholder = "https://api.example.com/path  or  {{BASE_URL}}/path"
	url.Focus()
	url.Width = 66
	url.TextStyle = brightStyle

	token := textinput.New()
	token.Placeholder = "leave empty to use env TOKEN"
	token.Width = 66
	token.EchoMode = textinput.EchoPassword
	token.EchoCharacter = '•'

	save := textinput.New()
	save.Placeholder = "name to save (leave empty to skip)"
	save.Width = 66

	body := textarea.New()
	body.Placeholder = "{\n  \"key\": \"value\"\n}"
	body.SetWidth(66)
	body.SetHeight(7)
	body.ShowLineNumbers = false

	return RequestFormModel{
		method:     method,
		urlInput:   url,
		tokenInput: token,
		saveInput:  save,
		bodyArea:   body,
		focused:    fieldURL,
	}
}

func (m RequestFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m RequestFormModel) Update(msg tea.Msg) (RequestFormModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+r":
			if !m.loading {
				m.loading = true
				m.err = nil
				m.savedName = ""
				return m, m.doSend()
			}

		case "esc":
			return m, navCmd(ScreenHome, "")

		case "tab":
			m.focused = (m.focused + 1) % fieldCount
			m = m.syncFocus()
			return m, textinput.Blink

		case "shift+tab":
			m.focused = (m.focused - 1 + fieldCount) % fieldCount
			m = m.syncFocus()
			return m, textinput.Blink
		}

	case responseMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}

		return m, func() tea.Msg {
			return NavigateMsg{To: ScreenResponse, Response: msg.resp}
		}
	}

	var cmd tea.Cmd
	switch m.focused {
	case fieldURL:
		m.urlInput, cmd = m.urlInput.Update(msg)
	case fieldToken:
		m.tokenInput, cmd = m.tokenInput.Update(msg)
	case fieldBody:
		m.bodyArea, cmd = m.bodyArea.Update(msg)
	case fieldSave:
		m.saveInput, cmd = m.saveInput.Update(msg)
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m RequestFormModel) syncFocus() RequestFormModel {
	m.urlInput.Blur()
	m.tokenInput.Blur()
	m.bodyArea.Blur()
	m.saveInput.Blur()

	switch m.focused {
	case fieldURL:
		m.urlInput.Focus()
	case fieldToken:
		m.tokenInput.Focus()
	case fieldBody:
		m.bodyArea.Focus()
	case fieldSave:
		m.saveInput.Focus()
	}

	return m
}

func (m RequestFormModel) View() string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString("  " + titleStyle.Render("ranpo") + "  " + methodBadge(m.method) + "\n")
	sb.WriteString("  " + dividerLine(70) + "\n\n")

	sb.WriteString(m.renderInput("URL", m.urlInput.View(), m.focused == fieldURL))
	sb.WriteString(m.renderInput("Token", m.tokenInput.View(), m.focused == fieldToken))
	sb.WriteString(m.renderBody())
	sb.WriteString(m.renderInput("Save as", m.saveInput.View(), m.focused == fieldSave))

	switch {
	case m.loading:
		sb.WriteString("  " + dimStyle.Render("sending...") + "\n")
	case m.err != nil:
		sb.WriteString("  " + errorStyle.Render("✗ "+m.err.Error()) + "\n")
	case m.savedName != "":
		sb.WriteString("  " + successStyle.Render(fmt.Sprintf("✓ saved as %q", m.savedName)) + "\n")
	}

	sb.WriteString("\n")
	sb.WriteString(helpBar("tab", "next field", "shift+tab", "prev", "ctrl+s", "send", "esc", "home"))
	return sb.String()
}

func (m RequestFormModel) renderInput(lbl, content string, focused bool) string {
	var l string
	var box string

	if focused {
		l = labelStyle.Render(lbl)
		box = focusedBorderStyle.Render(content)
	} else {
		l = mutedStyle.Render(lbl)
		box = blurredBorderStyle.Render(content)
	}

	return "  " + l + "\n" + lipgloss.NewStyle().PaddingLeft(2).Render(box) + "\n\n"
}

func (m RequestFormModel) renderBody() string {
	lbl := mutedStyle.Render("Body (JSON)")
	box := blurredBorderStyle.Render(m.bodyArea.View())

	if m.focused == fieldBody {
		lbl = labelStyle.Render("Body (JSON)")
		box = focusedBorderStyle.Render(m.bodyArea.View())
	}

	return "  " + lbl + "\n" + lipgloss.NewStyle().PaddingLeft(2).Render(box) + "\n\n"
}

func (m RequestFormModel) doSend() tea.Cmd {
	return func() tea.Msg {
		cfg, _ := config.Load()
		env, _ := store.LoadEnv(cfg.ActiveEnv)

		// Auth resolution priority: inline token > env TOKEN > config default
		auth := models.AuthConfig{Type: "none"}

		if token := m.tokenInput.Value(); token != "" {
			auth = models.AuthConfig{Type: "bearer", Token: token}
		} else if env != nil {
			if token, ok := env.Variables["TOKEN"]; ok && token != "" {
				auth = models.AuthConfig{Type: "bearer", Token: token}
			}
		} else if cfg.DefaultAuth.Token != "" {
			auth = models.AuthConfig{Type: "bearer", Token: cfg.DefaultAuth.Token}
		}

		req := &models.Request{
			Method:  m.method,
			URL:     m.urlInput.Value(),
			Body:    m.bodyArea.Value(),
			Auth:    auth,
			Headers: map[string]string{},
		}

		// Persist if name given
		if name := strings.TrimSpace(m.saveInput.Value()); name != "" {
			req.Name = name

			col, _ := store.LoadCollection("default")
			if col == nil {
				col = &models.Collection{Name: "default"}
			}

			replaced := false
			for i, r := range col.Requests {
				if r.Name == name {
					col.Requests[i] = *req
					replaced = true
					break
				}
			}

			if !replaced {
				col.Requests = append(col.Requests, *req)
			}

			_ = store.SaveCollection(col)
		}

		resp, err := client.Execute(req, env)
		return responseMsg{resp: resp, err: err}
	}
}
