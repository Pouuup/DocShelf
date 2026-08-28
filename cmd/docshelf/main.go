package main

import (
	"fmt"
	"os"

	"github.com/Poup-puoP/DocShelf/internal/app"
	"github.com/Poup-puoP/DocShelf/internal/parser"
	"github.com/Poup-puoP/DocShelf/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	url := "https://habr.com/ru/articles/881014/"
	pathStorage := "data"
	err := app.ImportDocument(url, pathStorage)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Document was successfully imported")
	}

	document, err := parser.ParseDocument(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := tui.NewModel(document)

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
