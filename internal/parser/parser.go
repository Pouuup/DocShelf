package parser

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Poup-puoP/DocShelf/internal/models"

	"golang.org/x/net/html"
)

func ParseDocument(url string) (models.Document, error) {
	htmlData, err := downloadHTML(url)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	title, err := extractTitle(htmlData)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	body, err := extractBody(htmlData)
	if err != nil {
		return models.Document{}, fmt.Errorf("%w", err)
	}

	sections := parseSections(body)

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

func extractTitle(htmlData []byte) (string, error) {
	if len(htmlData) == 0 {
		return "", fmt.Errorf("HTML is empty")
	}

	reader := bytes.NewReader(htmlData)

	root, err := html.Parse(reader)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	title := findTitle(root)
	if title == nil {
		return "", fmt.Errorf("Title not found")
	}

	if title.FirstChild == nil {
		return "", fmt.Errorf("Title is empty")
	}
	if title.FirstChild.Data == "" {
		return "", fmt.Errorf("Title is empty")
	}

	return title.FirstChild.Data, nil

}

func findTitle(node *html.Node) *html.Node {
	if node == nil {
		return node
	}
	if node.Type == html.ElementNode {
		if node.Data == "title" {
			return node
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {

		find := findTitle(child)
		if find != nil {
			return find
		}

	}

	return nil
}

func extractBody(htmlData []byte) (*html.Node, error) {
	if len(htmlData) == 0 {
		return nil, fmt.Errorf("HTML is empty")
	}

	reader := bytes.NewReader(htmlData)

	root, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	body := findBody(root)
	if body == nil {
		return nil, fmt.Errorf("Body not found")
	}

	return body, nil

}

func findBody(node *html.Node) *html.Node {
	if node == nil {
		return node
	}
	if node.Type == html.ElementNode {
		if node.Data == "body" {
			return node
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		find := findBody(child)
		if find != nil {
			return find
		}
	}

	return nil
}

func buildSections(node *html.Node, section *[]models.Section, currentSections *models.Section) {

	if node == nil {
		return
	}

	if node.Type == html.ElementNode {
		if node.Data == "h2" {
			title := extractText(node)
			if currentSections.Title != "" {
				*section = append(*section, *currentSections)
			}
			*currentSections = models.Section{}
			currentSections.Title = title
		}

		if node.Data == "p" {
			text := extractText(node)
			currentSections.Text += text + "\n"
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		buildSections(child, section, currentSections)

	}
}

func parseSections(node *html.Node) []models.Section {
	sections := []models.Section{}
	currentSection := models.Section{}

	buildSections(node, &sections, &currentSection)
	if currentSection.Title != "" {
		sections = append(sections, currentSection)
	}
	return sections
}

func extractText(text *html.Node) string {
	if text == nil {
		return ""
	}

	if text.Type == html.TextNode {
		return text.Data
	}

	totalText := ""

	for child := text.FirstChild; child != nil; child = child.NextSibling {
		childText := extractText(child)
		totalText += childText
	}

	return totalText
}
