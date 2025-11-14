package button

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PressedMsg struct {
	Key tea.KeyType
}

type Model struct {
	Key         tea.KeyType
	Active      bool
	title       string
	width       int
	heightLines int
	windowWidth int
}

func NewModel(title string, key tea.KeyType, width, height int) *Model {
	return &Model{
		title:       title,
		Key:         key,
		width:       width,
		heightLines: height,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case m.Key:
			return m, func() tea.Msg {
				return PressedMsg{Key: m.Key}
			}
		}
	}

	return m, nil
}

func (m *Model) View() string {
	border := 1
	var btnWidth int
	if m.width > 0 {
		btnWidth = m.width - 2*border
	} else {
		btnWidth = m.windowWidth - 2*border
	}
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(btnWidth).
		Height(m.heightLines).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)

	if m.Active {
		style = style.
			BorderForeground(lipgloss.Color("#00FF88")). // зелёная рамка
			Bold(true)
	} else {
		style = style.
			BorderForeground(lipgloss.Color("#666666"))
	}

	return style.Render(m.title)
}
