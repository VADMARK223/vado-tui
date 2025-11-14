package menu

import (
	"strings"
	"vado-tui/component/button"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	gap         = 2
	buttonWidth = 17
)

type SelectMsg struct {
	Key tea.KeyType
}

type area struct {
	start int
	end   int
	key   tea.KeyType
}

type Model struct {
	terminalHeight int
	btn1           *button.Model
	btn2           *button.Model
	btn3           *button.Model
	btn10          *button.Model
	areas          []area
}

func NewModel() *Model {
	return &Model{
		btn1:  button.NewModel("F1 Help", tea.KeyF1, buttonWidth, 1),
		btn2:  button.NewModel("F2 Chat", tea.KeyF2, buttonWidth, 1),
		btn3:  button.NewModel("F3 Settings", tea.KeyF3, buttonWidth, 1),
		btn10: button.NewModel("F10 Quit", tea.KeyF10, buttonWidth, 1),
		areas: []area{
			{start: 0, end: buttonWidth, key: tea.KeyF1},
			{start: buttonWidth + gap, end: buttonWidth + gap + buttonWidth, key: tea.KeyF2},
			{start: 2 * (buttonWidth + gap), end: 2*(buttonWidth+gap) + buttonWidth, key: tea.KeyF3},
			{start: 3 * (buttonWidth + gap), end: 3*(buttonWidth+gap) + buttonWidth, key: tea.KeyF10},
		},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalHeight = msg.Height

	case tea.MouseMsg:
		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft {
			if msg.Y >= m.terminalHeight-5 && msg.Y <= m.terminalHeight-gap {
				x := msg.X
				for _, a := range m.areas {
					if x >= a.start && x <= a.end {
						m.activate(a.key)
						return m, func() tea.Msg {
							return SelectMsg{Key: a.key}
						}
					}
				}
			}
		}
	}

	switch msg := msg.(type) {
	case button.PressedMsg:
		m.activate(msg.Key)
		return m, func() tea.Msg {
			return SelectMsg{Key: msg.Key}
		}
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.btn1, cmd = m.btn1.Update(msg)
	cmds = append(cmds, cmd)
	m.btn2, cmd = m.btn2.Update(msg)
	cmds = append(cmds, cmd)
	m.btn3, cmd = m.btn3.Update(msg)
	cmds = append(cmds, cmd)
	m.btn10, cmd = m.btn10.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	return joinManyHorizontally("  ", m.btn1.View(), m.btn2.View(), m.btn3.View(), m.btn10.View()) + "\n"
}

func (m *Model) activate(key tea.KeyType) {
	m.btn1.Active = m.btn1.Key == key
	m.btn2.Active = m.btn2.Key == key
	m.btn3.Active = m.btn3.Key == key
}

func joinManyHorizontally(sep string, blocks ...string) string {
	if len(blocks) == 0 {
		return ""
	}
	out := blocks[0]
	for _, b := range blocks[1:] {
		out = joinHorizontal(out, b, sep)
	}
	return out
}

func joinHorizontal(a, b string, sep string) string {
	linesA := strings.Split(a, "\n")
	linesB := strings.Split(b, "\n")

	maxLines := len(linesA)
	if len(linesB) > maxLines {
		maxLines = len(linesB)
	}

	out := make([]string, maxLines)
	for i := 0; i < maxLines; i++ {
		var left, right string
		if i < len(linesA) {
			left = linesA[i]
		}
		if i < len(linesB) {
			right = linesB[i]
		}
		out[i] = left + sep + right
	}
	return strings.Join(out, "\n")
}
