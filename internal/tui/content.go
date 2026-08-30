package tui

func (m model) contentLines() []string {
	section := m.document.Sections[m.currentSection]

	var lines []string
	for _, block := range section.Blocks {
		text := WrapText(block.Text, m.width)
		lines = append(lines, text...)
	}

	return lines
}
