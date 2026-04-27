package tui

import (
	"fmt"
	"strings"

	"github.com/awiipp/ranpo/internal/config"
	"github.com/awiipp/ranpo/internal/store"
	"github.com/awiipp/ranpo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
)

type EnvModel struct {
	envNames  []string
	cursor    int
	envs      []*models.Environment
	activeEnv string
	status    string
}

func NewEnvModel() EnvModel {
	return EnvModel{}
}

func (m EnvModel) Init() tea.Cmd {
	return m.loadEnvs()
}

func (m EnvModel) Update(msg tea.Msg) (EnvModel, tea.Cmd) {
	switch msg := msg.(type) {

	case envsLoadedMsg:
		m.envNames = msg.names
		m.envs = msg.envs
		m.activeEnv = msg.active

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, func() tea.Msg { return NavigateMsg{To: ScreenHome} }

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.envNames)-1 {
				m.cursor++
			}

		case "enter", " ":
			if len(m.envNames) > 0 {
				name := m.envNames[m.cursor]
				cfg, _ := config.Load()
				cfg.ActiveEnv = name
				_ = config.Save(cfg)
				m.activeEnv = name
				m.status = fmt.Sprintf("switched to %q", name)
			}
		}
	}

	return m, nil
}

func (m EnvModel) View() string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString("  " + titleStyle.Render("ranpo") + "  " + dimStyle.Render("environments") + "\n")
	sb.WriteString("  " + dimmerStyle.Render(strings.Repeat("─", 48)) + "\n\n")

	if len(m.envNames) == 0 {
		sb.WriteString("  " + dimStyle.Render("no environments yet.") + "\n")
		sb.WriteString("  " + dimStyle.Render("run: ranpo env set <name> KEY VALUE") + "\n")
	} else {
		for i, name := range m.envNames {
			isActive := name == m.activeEnv
			isCursor := i == m.cursor

			bullet := "  "
			nameStr := normalItemStyle.Render(name)

			if isCursor {
				bullet = selectedItemStyle.Render("❯ ")
				nameStr = selectedItemStyle.Render(name)
			}

			activeMark := ""
			if isActive {
				activeMark = "  " + successStyle.Render("● active")
			}

			sb.WriteString("  " + bullet + nameStr + activeMark + "\n")

			// Show variables for active/cursor env
			if (isCursor || isActive) && i < len(m.envs) && m.envs[i] != nil {
				for k, v := range m.envs[i].Variables {
					display := v
					if strings.Contains(strings.ToLower(k), "token") ||
						strings.Contains(strings.ToLower(k), "secret") ||
						strings.Contains(strings.ToLower(k), "pass") {
						display = "••••••••••••"
					}

					sb.WriteString("      " +
						dimStyle.Render(fmt.Sprintf("%-12s", k)) +
						dimmerStyle.Render("= ") +
						mutedStyle.Render(display) + "\n")
				}
			}

			sb.WriteString("\n")
		}
	}

	if m.status != "" {
		sb.WriteString("  " + successStyle.Render("✓ "+m.status) + "\n\n")
	}

	sb.WriteString(helpBar(
		"↑↓/jk", "navigate",
		"enter", "switch active env",
		"esc", "back",
	))
	sb.WriteString("\n  " + dimStyle.Render("tip: use  ranpo env set <env> KEY VALUE  to add variables") + "\n")

	return sb.String()
}

type envsLoadedMsg struct {
	names  []string
	envs   []*models.Environment
	active string
}

func (m EnvModel) loadEnvs() tea.Cmd {
	return func() tea.Msg {
		names, _ := store.ListEnvs()
		cfg, _ := config.Load()

		envs := make([]*models.Environment, len(names))
		for i, name := range names {
			env, _ := store.LoadEnv(name)
			envs[i] = env
		}

		return envsLoadedMsg{names: names, envs: envs, active: cfg.ActiveEnv}
	}
}
