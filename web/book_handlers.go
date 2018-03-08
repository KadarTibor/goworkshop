package web

import (
	"net/http"
	"goworkshop/model"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request){
	WriteJson(w, model.Books)
	w.WriteHeader(http.StatusOK)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, book)

	if err != nil {
		fmt.Fprintf(w, "Failed to read body")
	} else {
		w.WriteHeader(http.StatusOK)
		model.Books.AddBook(book)
		WriteJson(w, book)
	}
}

func GetBookByUUID(w http.ResponseWriter, r *http.Request) {
	bookUuid := mux.Vars(r)["uuid"]

	book, err := model.Books.GetBookByUuid(bookUuid)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	} else {
		WriteJson(w, book)
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteBookByUUID(w http.ResponseWriter, r *http.Request){
	bookUuid := mux.Vars(r)["uuid"]
	err := model.Books.DeleteBookWithUuid(bookUuid)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	} else {
		WriteJson(w, model.Books)
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var book model.Book
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, book)
	if err != nil {
		fmt.Fprintf(w, "Failed to read body %s", err)
		return
	} else {
		book, err = model.Books.UpdateBookWithUuid(book)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		} else {
			WriteJson(w, book)
			w.WriteHeader(http.StatusOK)
		}
	}
}