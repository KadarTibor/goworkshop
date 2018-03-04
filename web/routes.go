package web

import "net/http"

const (
	authorsBaseUrl = "/authors"
	authorsByUUID = authorsBaseUrl + "/{uuid}"
	booksBaseUrl = "/books"
	booksByUUID = booksBaseUrl + "/{uuid}"
)

type Route struct {
	Method string
	Pattern string
	ExpectedCode int
	HandlerFunc http.HandlerFunc
}

