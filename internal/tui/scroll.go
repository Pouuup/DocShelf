package tui

import "strings"

func (m model) maxPosition() int {
	lines := strings.Split(
		m.document.Sections[m.currentSection].Text,
		"\n",
	)
	pageSize := m.height - 4
	if len(lines) <= pageSize {
		return 0
	}

	return len(lines) - pageSize
}
