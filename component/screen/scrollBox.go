package screen

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ScrollBox struct {
	lines  []string
	offset int
	width  int
	height int
}

func NewScrollBox(text string, width, height int) *ScrollBox {
	return &ScrollBox{
		lines:  strings.Split(text, "\n"),
		offset: 0,
		width:  width,
		height: height,
	}
}

func (s *ScrollBox) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	// клавиши
	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if s.offset > 0 {
				s.offset--
			}

		case "down", "j":
			if s.offset < len(s.lines)-s.height {
				s.offset++
			}

		case "pgup":
			s.offset -= s.height
			if s.offset < 0 {
				s.offset = 0
			}

		case "pgdown":
			s.offset += s.height
			if s.offset > len(s.lines)-s.height {
				s.offset = len(s.lines) - s.height
			}
		}

	// колесо мыши
	case tea.MouseMsg:
		switch msg.Button {
		case tea.MouseButtonWheelUp:
			if s.offset > 0 {
				s.offset--
			}
		case tea.MouseButtonWheelDown:
			if s.offset < len(s.lines)-s.height {
				s.offset++
			}
		}
	}

	return nil
}

func (s *ScrollBox) View() string {
	end := s.offset + s.height
	if end > len(s.lines) {
		end = len(s.lines)
	}

	visible := s.lines[s.offset:end]

	style := lipgloss.NewStyle().
		Width(s.width).
		Height(s.height)

	return style.Render(strings.Join(visible, "\n"))
}

func (s *ScrollBox) Resize(width, height int) {
	s.width = width
	s.height = height
}
