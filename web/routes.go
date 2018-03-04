package web

import "net/http"

const (
	authorsBaseUrl = "/authors"
	authorsByUuidUrl  = authorsBaseUrl + "/{uuid}"
	booksBaseUrl   = "/books"
	booksByUuidUrl    = booksBaseUrl + "/{uuid}"
)

type Route struct {
	Method       string
	Pattern      string
	ExpectedCode int
	HandlerFunc  http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	//book_handlers
	Route{
		Method:      "GET",
		Pattern:     booksBaseUrl,
		HandlerFunc: GetAllBooks,
	},
	Route{
		Method:      "POST",
		Pattern:     booksBaseUrl,
		HandlerFunc: AddBook,
	},
	Route{
		Method:      "GET",
		Pattern:     booksByUuidUrl,
		HandlerFunc: GetBookByUUID,
	},
	Route{
		Method:      "DELETE",
		Pattern:     booksByUuidUrl,
		HandlerFunc: DeleteBookByUUID,
	},
	Route{
		Method:      "PUT",
		Pattern:     booksByUuidUrl,
		HandlerFunc: UpdateBook,
	},

	//author_handlers
	Route{
		Method:      "GET",
		Pattern:     authorsBaseUrl,
		HandlerFunc: GetAllAuthors,
	},
	Route{
		Method:      "POST",
		Pattern:     authorsBaseUrl,
		HandlerFunc: AddAuthor,
	},
	Route{
		Method:      "GET",
		Pattern:     authorsByUuidUrl,
		HandlerFunc: GetAuthorByUUID,
	},
	Route{
		Method:      "DELETE",
		Pattern:     authorsByUuidUrl,
		HandlerFunc: DeleteAuthorByUUID,
	},
	Route{
		Method:      "PUT",
		Pattern:     authorsByUuidUrl,
		HandlerFunc: UpdateAuthor,
	},
}
