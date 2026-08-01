package parser

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

func downloadHTML(url string) ([]byte, error) {
	var html []byte
	if url == "" {
		return html, fmt.Errorf("error: URL is empty")
	}

	resp, err := http.Get(url)
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

		val := findTitle(child)
		if val != nil {
			return val
		}

	}

	return nil
}

