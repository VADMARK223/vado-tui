package screen

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Type uint

const (
	ScreenDefaul Type = iota
	ScreenHelp
	ScreenChat
	ScreenSettings
)

type Model struct {
	width  int
	height int

	menuHeight int

	active Type
}

func NewModel() *Model {
	return &Model{menuHeight: 6, active: ScreenDefaul}
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
	}

	return m, nil
}

func (m *Model) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.width).
		Height(m.height)

	switch m.active {
	case ScreenHelp:
		return style.Render("HELP SCREEN:\n\nВот тут хелп текст...")
	case ScreenChat:
		return style.Render("Chat:\n\nВот тут хелп текст...")
	case ScreenSettings:
		return style.Render("Settings:\n\nВот тут хелп текст...")
	case ScreenDefaul:
		fallthrough
	default:
		return style.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Render("Vado TUI")
	}
}
