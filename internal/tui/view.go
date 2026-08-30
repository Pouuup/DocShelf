package tui

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func (m model) View() string {
	var view strings.Builder

	totalSections := len(m.document.Sections) // Total number of sections

	// If there are no sections, we return the message
	if totalSections == 0 {
		return "No sections\n"
	}

	// Display [Current Section/Total Section]
	progresSection := "[" + strconv.Itoa(m.currentSection+1) + "/" + strconv.Itoa(totalSections) + "]"
	title := m.document.Sections[m.currentSection].Title + "                                                                " + progresSection

	lines := m.contentLines()

	pageSize := m.height - 5 // How many lines can fit on the screen

	if pageSize <= 0 {
		pageSize = 10
	}

	end := m.position + pageSize // End of page display

	if end > len(lines) {
		end = len(lines)
	}

	// Adding lines in 'view'
	for i := m.position; i < end; i++ {
		view.WriteString(lines[i])
		view.WriteString("\n")
	}

	// Adding the missing lines to display Key Hints
	visibleLines := end - m.position
	for i := 0; i < (pageSize - visibleLines); i++ {
		view.WriteString("\n")
	}

	// The final display
	partOne := title + "\n" + "─────────────────────────────────────────────────────────────────────────────────────" + "\n" + view.String()
	partTwo := "\n" + "─────────────────────────────────────────────────────────────────────────────────────" + "\n" + "↑↓: scroll;   ←→: section;   q: quit"
	total := partOne + partTwo
	return total
}

func WrapText(text string, width int) []string {
	var lines []string
	var current strings.Builder

	words := strings.Fields(text)

	for _, word := range words {
		currentLen := utf8.RuneCountInString(current.String())
		wordLen := utf8.RuneCountInString(word)

		var newLen int

		if currentLen == 0 {
			newLen = currentLen + wordLen
		} else {
			newLen = currentLen + 1 + wordLen
		}

		if newLen <= width {
			if current.Len() > 0 {
				current.WriteString(" ")
			}
			current.WriteString(word)

		} else {
			lines = append(lines, current.String())
			current.Reset()
			current.WriteString(word)
		}

	}
	if current.Len() > 0 {
		lines = append(lines, current.String())
	}

	return lines
}
