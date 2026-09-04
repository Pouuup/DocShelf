package app

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Pouuup/DocShelf/internal/models"
	"github.com/Pouuup/DocShelf/internal/parser"
	"github.com/Pouuup/DocShelf/internal/storage"
)

func ImportDocument(url string, pathStorage string) (models.Document, error) {
	document, err := parser.ParseDocument(url)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	titleFile := sanitizeFilename(document.Title)
	pathDir := filepath.Join(pathStorage, titleFile)

	exist, err := storage.ExistsByURL(pathStorage, document.SourceURL)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	if exist {
		err = storage.Update(&document, pathStorage, pathDir)
		if err != nil {
			return models.Document{}, fmt.Errorf("%w", err)
		}
	} else {
		err = storage.Save(&document, pathStorage, pathDir)
		if err != nil {
			return models.Document{}, fmt.Errorf("%w", err)
		}
	}

	return document, nil
}

func sanitizeFilename(name string) string {
	replacer := strings.NewReplacer(
		"/", "-",
		"\\", "-",
		":", "-",
		"*", "_",
		"?", "_",
		"\"", "'",
		"<", "(",
		">", ")",
		"|", "-",
	)
	return strings.TrimSpace(replacer.Replace(name))
}
