package tui

import "github.com/Poup-puoP/DocShelf/internal/models"

type model struct {
    document       models.Document
    currentSection int
    position       int
}