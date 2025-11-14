package tabs

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	cursor int
	width  int
	height int
}

func NewModel() *Model {
	return &Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < 1 {
				m.cursor++
			}

		case "enter":
			if m.cursor == 0 {
				fmt.Println("🚀 Нажата первая кнопка")
			} else if m.cursor == 1 {
				fmt.Println("🔥 Нажата вторая кнопка")
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	buttonStyle := lipgloss.NewStyle().
		//Padding(0, 2).
		Border(lipgloss.RoundedBorder()).
		Width(30)

	focusedStyle := buttonStyle.Copy().
		BorderForeground(lipgloss.Color("#00BFFF")).
		Bold(true)

	buttons := []string{"Кнопка 1", "Кнопка 2"}

	var out string
	for i, label := range buttons {
		if i == m.cursor {
			out += focusedStyle.Render(label) + "\n"
		} else {
			out += buttonStyle.Render(label) + "\n"
		}
	}
	return lipgloss.PlaceHorizontal(m.width, lipgloss.Right, out)
}
