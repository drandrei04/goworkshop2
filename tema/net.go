package main

import (
	"fmt"
)

//"github.com/satori/go.uuid"

type person struct {
	name string
	age  int
}

type AuthorDto struct {
	UUID      string
	FirstName string
	LastName  string
	Birthday  string
	Death     string
}

type BookDto struct {
	UUID        string
	Title       string
	ReleaseDate string
	Author      AuthorDto
}

func main() {


	b := BookDto{Title: "Capra", ReleaseDate: "1900 toamna"}
	b.Author.FirstName = "Ion"
	b.Author.LastName = "Creanga"
	b.Author.Death = "1900iarna"
	b.Author.Birthday = "1900primavara"
	fmt.Println(b)

	