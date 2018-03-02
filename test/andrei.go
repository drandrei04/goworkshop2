package main

import (
	"testing"
)

func Main(t *testing.T) {

	type AuthorDto struct {
		FirstName string
		LastName  string
		Birthday  string
		Death     string
	}

	type BookDto struct {
		Title       string
		ReleaseDate string
		Author      AuthorDto
	}

}
