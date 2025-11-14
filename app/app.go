package app

import (
	"fmt"
	"vado-tui/component/menu"
	"vado-tui/component/screen"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	scr           *screen.Model
	menu          *menu.Model
	currentScreen uint
}

func NewModel() *Model {
	return &Model{
		scr:  screen.NewModel(),
		menu: menu.NewModel(),
	}
}

func (m *Model) Init() tea.Cmd {
	//m.scr.SetScreen(screen.Help) // TODO Delete after
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	_, cmd := m.scr.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	switch msg := msg.(type) {

	case menu.SelectMsg:
		if msg.Key == tea.KeyF10 {
			return m, tea.Quit
		}
		m.onMenuSelect(msg.Key)
	}

	_, cmd = m.menu.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	return fmt.Sprintf("%s\n%s", m.scr.View(), m.menu.View())
}

func (m *Model) onMenuSelect(key tea.KeyType) {
	switch key {
	case tea.KeyF1:
		m.scr.SetScreen(screen.Help)
	case tea.KeyF2:
		m.scr.SetScreen(screen.Chat)
	case tea.KeyF3:
		m.scr.SetScreen(screen.Settings)
	default:
		panic("invalid key")
	}
}
