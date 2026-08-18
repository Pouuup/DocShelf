package main

import (
	"fmt"

	"github.com/Poup-puoP/DocShelf/internal/app"
)

func main() {
	url := ""
	pathStorage := "data"
	err := app.ImportDocument(url, pathStorage)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Document was successfully imported")
	}
}
