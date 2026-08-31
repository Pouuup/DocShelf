package tui

import (
	"github.com/Poup-puoP/DocShelf/internal/models"
)

func (m model) contentLines() []string {
	section := m.document.Sections[m.currentSection]

	var lines []string
	for _, block := range section.Blocks {
		switch block.Type {
		case models.BlockParagraph:
			text := WrapText(block.Text, m.width)
			lines = append(lines, text...)
		case models.BlockCode:
			codeLines := renderCode(block.Text)
			lines = append(lines, codeLines...)
		}
	}

	return lines
}
