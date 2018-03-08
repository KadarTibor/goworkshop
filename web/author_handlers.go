package web

import (
	"net/http"
	"goworkshop/model"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, model.Authors)
	w.WriteHeader(http.StatusOK)
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, author)

	if err != nil {
		fmt.Fprintf(w, "Failed to read body")
	} else {
		w.WriteHeader(http.StatusOK)
		model.Authors.AddAuthor(author)
		WriteJson(w, author)
	}
}

func GetAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	authorUuid := mux.Vars(r)["uuid"]

	author, err := model.Authors.GetAuthorByUuid(authorUuid)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	} else {
		WriteJson(w, author)
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	authorUuid := mux.Vars(r)["uuid"]
	err := model.Authors.DeleteAuthorWithUuid(authorUuid)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	} else {
		WriteJson(w, model.Authors)
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, author)
	if err != nil {
		fmt.Fprintf(w, "Failed to read body %s", err)
		return
	} else {
		author, err = model.Authors.UpdateAuthorWithUuid(author)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		} else {
			WriteJson(w, author)
			w.WriteHeader(http.StatusOK)
		}
	}
}
