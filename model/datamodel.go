package model

import (
	"fmt"
)

//AuthorDto - Author structure
type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

//BookDto - Book structure
type BookDto struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	NoPages     int `json:"noPages"`
	ReleaseDate string `json:"releaseDate"`
	Author      AuthorDto `json:"author"`
}


func (b BookDto) String() string{
	return fmt.Sprintf("BookDto{ UUID = %s , Title = %s, NoPages = %d, ReleaseDate = %s, Author = %s}", b.UUID,
		b.Title,b.NoPages, b.ReleaseDate, b.Author)
}

func (a AuthorDto) String() string{
	return fmt.Sprintf( "AuthorDto{ UUID = %s, FirstName = %s, LastName = %s, Birthday = %s, Death = %s}",
		a.UUID,a.FirstName, a.LastName, a.Birthday, a.Death)
}

var authors []AuthorDto
var books []BookDto
