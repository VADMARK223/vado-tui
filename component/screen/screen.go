package screen

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Type uint

const (
	Default Type = iota
	Help
	Chat
	Settings
)

type Model struct {
	scroll     *ScrollBox
	active     Type
	width      int
	height     int
	menuHeight int
}

func NewModel() *Model {
	return &Model{menuHeight: 6, active: Default}
}

func (m *Model) SetScreen(active Type) {
	m.active = active
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width - 2
		m.height = msg.Height - m.menuHeight

		if m.scroll != nil {
			m.scroll.Resize(m.width-4, m.height)
		}
	}

	if m.active == Help && m.scroll != nil {
		m.scroll.Update(msg)
	}
	return m, nil
}

func (m *Model) View() string {
	outer := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.width).
		Height(m.height)

	switch m.active {
	case Help:
		helpText := `HELP SCREEN:
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20`
		if m.scroll == nil {
			m.scroll = NewScrollBox(helpText, m.width-4, m.height)
		}

		return outer.Render(m.scroll.View())
	case Chat:
		return outer.Render("CHAT SCREEN")
	case Settings:
		return outer.Render("SETTINGS SCREEN")
	case Default:
		fallthrough
	default:
		return outer.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Render("Vado TUI")
	}
}
