package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Pouuup/DocShelf/internal/app"
	"github.com/Pouuup/DocShelf/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var url string
	pathStorage := "data"

	fmt.Println("Enter the link:")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		url = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	document, err := app.ImportDocument(url, pathStorage)
	if err != nil {
		fmt.Println(err)
		return
	} /*else {
		fmt.Println("Document was successfully imported")
	}*/

	m := tui.NewModel(document)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
