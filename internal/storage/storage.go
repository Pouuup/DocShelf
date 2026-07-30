package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Poup-puoP/DocShelf/internal/models"
)

func SaveData(document *models.Document, pathStorage string, pathDirDoc string) error {
	if document == nil {
		return fmt.Errorf("error: the document is empty")
	}
	exist, err := ExistsByURL(pathStorage, document.Source_URL)
	if err != nil {
		return fmt.Errorf("error:%w", err)
	}

	if exist {
		return fmt.Errorf("error: The document exists")
	}

	err = os.MkdirAll(pathDirDoc, 0755)
	if err != nil {
		return fmt.Errorf("error:%w", err)
	}
	pathDocument := filepath.Join(pathDirDoc, "document.json")
	file, err := os.Create(pathDocument)
	if err != nil {
		return fmt.Errorf("error:%w", err)
	}
	defer file.Close()

	b, err := json.MarshalIndent(document, "", "	")
	if err != nil {
		return fmt.Errorf("error:%w", err)
	}

	_, err = file.Write(b)
	if err != nil {
		return fmt.Errorf("error:%w", err)
	}

	return nil
}

var ErrFoundURL = errors.New("URL match found")

func ExistsByURL(pathDir string, url string) (bool, error) {

	err := filepath.WalkDir(pathDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error:%w", err)
		}

		if !d.IsDir() && filepath.Ext(path) == ".json" {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("error:%w", err)
			}

			var doc struct {
				URL string `json:"Source_URL"`
			}

			err = json.NewDecoder(file).Decode(&doc)
			file.Close()

			if err != nil {
				return fmt.Errorf("error:%w", err)
			}

			if doc.URL == url {
				return ErrFoundURL
			}
		}

		return nil
	})

	if errors.Is(err, ErrFoundURL) {
		return true, nil
	}

	return false, err
}

func Load(pathDocument string) (models.Document, error) {
	var document models.Document

	if pathDocument == "" {
		return document, fmt.Errorf("error: Path is empty")
	}

	file, err := os.Open(pathDocument)
	if err != nil {
		return document, fmt.Errorf("error:%w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&document)
	if err != nil {
		return document, fmt.Errorf("error:%w", err)
	}
	return document, nil

}

var ErrNotDirectory = errors.New("path is not a directory")

func Delete(pathDir string) error {
	if pathDir == "" {
		return fmt.Errorf("error: Path is empty")
	}

	info, err := os.Stat(pathDir)
	if err != nil {
		return fmt.Errorf("error:%w", err)
	}

	if !info.IsDir() {
		return ErrNotDirectory
	}

	err = os.RemoveAll(pathDir)
	if err != nil {
		return fmt.Errorf("error:%w", err)
	}

	return nil
}

func List(pathDir string) ([]models.Document, error) {
	var documents []models.Document

	if pathDir == "" {
		return documents, fmt.Errorf("error: Path is empty")
	}

	info, err := os.Stat(pathDir)
	if err != nil {
		return documents, fmt.Errorf("error:%w", err)
	}

	if !info.IsDir() {
		return documents, ErrNotDirectory
	}

	err = filepath.WalkDir(pathDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error:%w", err)
		}

		if !d.IsDir() && filepath.Ext(path) == ".json" {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("error:%w", err)
			}
			defer file.Close()

			var doc models.Document

			err = json.NewDecoder(file).Decode(&doc)
			if err != nil {
				return fmt.Errorf("error:%w", err)
			}

			documents = append(documents, doc)
		}

		return nil
	})

	if err != nil {
		return documents, fmt.Errorf("error:%w", err)
	}

	return documents, nil

}
