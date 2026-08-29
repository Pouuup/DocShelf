package tui

import "github.com/Poup-puoP/DocShelf/internal/models"

type model struct {
	document       models.Document // Current document
	currentSection int             // Current Section
	position       int             // Cursor position
	width          int             // Terminal width
	height         int             // Terminal height
}

func NewModel(document models.Document) model {
	return model{
		document: document,
	}
}
