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

type Page struct {
	Content string
	Scroll  *ScrollBox
}

type Model struct {
	pages      map[Type]*Page
	active     Type
	width      int
	height     int
	menuHeight int
}

func NewModel() *Model {
	helpText := `HELP
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
	settingsText := `SETTINGS
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
20
21
22
23
24`
	m := &Model{menuHeight: 6, active: Default, pages: make(map[Type]*Page)}
	m.pages[Help] = &Page{Content: helpText}
	m.pages[Chat] = &Page{Content: "C"}
	m.pages[Settings] = &Page{Content: settingsText}
	m.pages[Default] = &Page{Content: "Vado TUI"}
	return m
}

func (m *Model) SetScreen(active Type) {
	m.active = active

	page := m.pages[m.active]
	if page.Scroll == nil && page.Content != "" {
		page.Scroll = NewScrollBox(page.Content, m.width-4, m.height)
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width - 2
		m.height = msg.Height - m.menuHeight

		for _, p := range m.pages {
			if p.Scroll != nil {
				p.Scroll.Resize(m.width-4, m.height)
			}
		}
	}

	page := m.pages[m.active]
	if page.Scroll != nil {
		page.Scroll.Update(msg)
	}

	return m, nil
}

func (m *Model) View() string {
	outer := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.width).
		Height(m.height)

	page := m.pages[m.active]

	if page.Scroll != nil {
		return outer.Render(page.Scroll.View())
	}

	return outer.
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Render(page.Content)

	/*if page.Scroll == nil && page.Content != "" {
		page.Scroll = NewScrollBox(page.Content, m.width-4, m.height)
	}

	if page.Scroll == nil {
		return outer.
			AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Render(page.Content)
	}

	return outer.Render(page.Scroll.View())*/
}
