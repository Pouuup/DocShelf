package parser

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Poup-puoP/DocShelf/internal/models"
	"github.com/PuerkitoBio/goquery"
)

func ParseDocument(url string) (models.Document, error) {
	htmlData, err := downloadHTML(url)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	documents, err := parseHTML(htmlData)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	title, err := extractTitle(documents)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	body, err := extractBody(documents)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	sections, err := parseSections(body)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	document := models.Document{
		Title:      title,
		SourceURL:  url,
		Sections:   sections,
		ImportedAt: time.Now(),
		UpdatedAt:  time.Now(),
	}

	return document, nil

}

func downloadHTML(url string) ([]byte, error) {
	var html []byte
	if url == "" {
		return html, fmt.Errorf("error: URL is empty")
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: unexpected HTTP status %d", url, resp.StatusCode)
	}

	html, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return html, nil
}

func extractTitle(document *goquery.Document) (string, error) {
	selection := document.Find("title").First()
	if selection.Length() == 0 {
		return "", fmt.Errorf("title not found")
	}

	return selection.Text(), nil
}

func extractBody(document *goquery.Document) (*goquery.Selection, error) {
	selection := document.Find("body").First()
	if selection.Length() == 0 {
		return nil, fmt.Errorf("body not found")
	}

	return selection, nil
}

func parseSections(body *goquery.Selection) ([]models.Section, error) {
	sections := []models.Section{}

	selection := body.Find("h2")
	selection.Each(func(i int, s *goquery.Selection) {
		section := models.Section{
			Title: s.Text(),
		}

		elements := s.NextUntil("h2")
		elements.Each(func(i int, p *goquery.Selection) {
			var block models.Block
			if p.Is("p") {
				block = models.Block{
					Type: models.BlockParagraph,
					Text: p.Text(),
				}
				section.Blocks = append(section.Blocks, block)
			} else if p.Is("pre") {
				block = models.Block{
					Type: models.BlockCode,
					Text: p.Text(),
				}
				section.Blocks = append(section.Blocks, block)
			}

		})

		sections = append(sections, section)

	})

	if selection.Length() == 0 {
		return nil, fmt.Errorf("no sections found")
	}

	return sections, nil

}

func parseHTML(htmlData []byte) (*goquery.Document, error) {
	reader := bytes.NewReader(htmlData)
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	return document, nil
}
