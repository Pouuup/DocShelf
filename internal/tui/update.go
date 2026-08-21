package tui

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "down":
			if m.position < m.maxPosition() {
				m.position++
			}
		case "up":
			if m.position > 0 {
				m.position--
			}
		case "right":
			if m.currentSection < len(m.document.Sections)-1 {
				m.currentSection++
				m.position = 0
			}

		case "left":
			if m.currentSection > 0 {
				m.currentSection--
				m.position = 0
			}
		}
	}
	return m, nil
}
