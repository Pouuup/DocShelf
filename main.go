package main

import (
	"fmt"

	"github.com/Poup-puoP/DocShelf/internal/app"
)

func main() {
	url := "https://go.dev/doc/"
	pathStorage := "data"
	err := app.ImportDocument(url, pathStorage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Document was successfully imported")
}
