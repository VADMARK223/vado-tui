package chat

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	messages []string
	textarea textarea.Model
	viewport viewport.Model
}

func NewModel() *Model {
	ta := textarea.New()
	ta.Placeholder = "Введите сообщение..."
	ta.Focus()

	vp := viewport.New(50, 15)
	vp.SetContent("")

	return &Model{
		messages: []string{},
		textarea: ta,
		viewport: vp,
	}
}

func (m *Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			text := strings.TrimSpace(m.textarea.Value())
			if text != "" {
				m.messages = append(m.messages, fmt.Sprintf("🧑 %s", text))
				m.viewport.SetContent(strings.Join(m.messages, "\n"))
				m.textarea.SetValue("")
				m.viewport.GotoBottom()
			}
			return m, nil
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00BFFF")).
		Render("💬 Мини-чат")

	return fmt.Sprintf("%s\n\n%s\n\n%s",
		title,
		m.viewport.View(),
		m.textarea.View(),
	)
}
