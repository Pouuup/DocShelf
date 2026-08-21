package tui

import "github.com/Poup-puoP/DocShelf/internal/models"

type model struct {
	document       models.Document
	currentSection int
	position       int
	width          int
	height         int
}

func NewModel(document models.Document) model {
	return model{
		document: document,
	}
}
