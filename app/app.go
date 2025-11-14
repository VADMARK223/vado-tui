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
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.scr.Update(msg)
	case menu.SelectMsg:
		m.onMenuSelect(msg.Key)
	}

	_, cmd := m.menu.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	return fmt.Sprintf("%s\n%s", m.scr.View(), m.menu.View())
}

func (m *Model) onMenuSelect(key tea.KeyType) {
	switch key {
	case tea.KeyF1:
		m.scr.SetScreen(screen.ScreenHelp)
	case tea.KeyF2:
		m.scr.SetScreen(screen.ScreenChat)
	case tea.KeyF3:
		m.scr.SetScreen(screen.ScreenSettings)
	default:
		panic("invalid key")
	}
}
