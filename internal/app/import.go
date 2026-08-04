package app

import (
	"fmt"
	"path/filepath"

	"github.com/Poup-puoP/DocShelf/internal/parser"
	"github.com/Poup-puoP/DocShelf/internal/storage"
)

func ImportDocument(url string, pathStorage string) error {
	document, err := parser.ParseDocument(url)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	pathDir := filepath.Join(pathStorage,document.Title)

	exist, err := storage.ExistsByURL(pathStorage, document.SourceURL)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if exist {
		err = storage.Update(&document, pathStorage, pathDir)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	} else {
		err = storage.Save(&document, pathStorage, pathDir)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	return nil 
}