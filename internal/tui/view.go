package tui

import (
	"strconv"
	"strings"
)

func (m model) View() string {
	var view strings.Builder

	totalSections := len(m.document.Sections)

	progresSection := "[" + strconv.Itoa(m.currentSection+1) + "/" + strconv.Itoa(totalSections) + "]"

	title := m.document.Sections[m.currentSection].Title + "                " + progresSection

	lines := strings.Split(
		m.document.Sections[m.currentSection].Text,
		"\n",
	)

	pageSize := m.height - 4

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

	partOne := title + "\n" + "────────────────────────────────────────────────────────" + "\n" + view.String()
	partTwo := "\n" + "────────────────────────────────────────────────────────" + "\n" + "↑↓: scroll;   ←→: section;   q: quit"
	total := partOne + partTwo
	return total
}
