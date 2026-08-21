package tui

import (
	"strings"
)

func (m model) View() string {
	var view strings.Builder

	lines := strings.Split(
		m.document.Sections[m.currentSection].Text,
		"\n",
	)

	pageSize := m.height

	if pageSize <= 0 {
		pageSize = 10
	}

	end := m.position + pageSize

	if end > len(lines) {
		end = len(lines)
	}

	for i := m.position; i < end; i++ {
		view.WriteString(lines[i])
		view.WriteString("\n")
	}

	return view.String()
}
