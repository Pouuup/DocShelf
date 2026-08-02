package parser

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

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

func extractHeadings(body *html.Node) []*html.Node {
	headings := make([]*html.Node, 0)
	if body == nil {
		return headings
	}

	if body.Type == html.ElementNode {
		if body.Data == "h1" || body.Data == "h2" || body.Data == "h3" || body.Data == "h4" || body.Data == "h5" || body.Data == "h6" {
			headings = append(headings, body)
		}
	}

	for child := body.FirstChild; child != nil; child = child.NextSibling {
		childHeadings := extractHeadings(child)
		headings = append(headings, childHeadings...)

	}

	return headings
}
