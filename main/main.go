package main

import (
	"fmt"
	"goworkshop/importer"
)

func main() {
	var authors = importer.ImportAuthors()
	fmt.Printf("Imported authors are: %s\n", authors)
	var books = importer.ImportBooks()
	fmt.Printf("Imported books are: %s\n", books)
}