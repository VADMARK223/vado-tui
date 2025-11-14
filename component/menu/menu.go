package menu

import (
	"strings"
	"vado-tui/component/button"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	buttonWidth = 17
)

type SelectMsg struct {
	Key tea.KeyType
}

type Model struct {
	btn1  *button.Model
	btn2  *button.Model
	btn3  *button.Model
	btn10 *button.Model
}

func NewModel() *Model {
	return &Model{
		btn1:  button.NewModel("F1 Help", tea.KeyF1, buttonWidth, 1),
		btn2:  button.NewModel("F2 Chat", tea.KeyF2, buttonWidth, 1),
		btn3:  button.NewModel("F3 Settings", tea.KeyF3, buttonWidth, 1),
		btn10: button.NewModel("F10 Quit", tea.KeyF10, buttonWidth, 1),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (*Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyF10:
			return m, tea.Quit
		default:
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

	switch msg := msg.(type) {
	case button.PressedMsg:
		m.activate(msg.Key)
		return m, func() tea.Msg {
			return SelectMsg{Key: msg.Key}
		}
	}

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
