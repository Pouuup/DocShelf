package tui

func (m model) maxPosition() int {
	lines := m.contentLines()

	pageSize := m.height - 5

	if pageSize <= 0 {
		pageSize = 10
	}

	if len(lines) <= pageSize {
		return 0
	}

	return len(lines) - pageSize
}
