package model

import "fmt"

//BookDto - Book structure
type Book struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	NoPages     int       `json:"noPages"`
	ReleaseDate string    `json:"releaseDate"`
	Author      Author `json:"author"`
}

func (b Book) String() string {
	return fmt.Sprintf("BookDto{ UUID = %s , Title = %s, NoPages = %d, ReleaseDate = %s, Author = %s}", b.UUID,
		b.Title, b.NoPages, b.ReleaseDate, b.Author)
}

func (a *BookList) GetBookByUuid(uuid string) (Book, error) {
	err := fmt.Errorf("Could not find any book with that UUID %s", uuid)

	for _, book := range *a {
		if book.UUID == uuid {
			return book, nil
		}
	}

	return Book{}, err
}

func (b *BookList) DeleteBookWithUuid(uuid string) (error) {
	var err = fmt.Errorf("Could not find any book with that uuid %s", uuid)
	var updatedBooks BookList
	for _, book := range *b {
		if book.UUID == uuid {
			err = nil
		} else {
			updatedBooks = append(updatedBooks, book)
		}
	}
	if err == nil {
		*b = updatedBooks
	}
	return err
}

func (b *BookList) UpdateBookWithUuid(updatedBook Book) (Book, error) {
	var err = fmt.Errorf("could not find book by uuid %s", updatedBook.UUID)
	var newBooks BookList
	for _, book := range *b {
		if book.UUID == updatedBook.UUID {
			newBooks = append(newBooks, updatedBook)
			err = nil
		} else {
			newBooks = append(newBooks, book)
		}
	}
	if err == nil {
		*b = newBooks
	}
	return updatedBook, err
}

func (a *BookList) AddBook(book Book) {
	*a = append(*a, book)
}

var Books BookList

type BookList []Book