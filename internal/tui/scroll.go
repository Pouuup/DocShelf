package tui

func (m model) maxPosition() int {
	lines := WrapText(
		m.document.Sections[m.currentSection].Text,
		m.width,
	)

	pageSize := m.height - 5

	if pageSize <= 0 {
		pageSize = 10
	}

	if len(lines) <= pageSize {
		return 0
	}

	return len(lines) - pageSize
}
