package main

import (
	"fmt"
	"goworkshop/model"
	"goworkshop/test"
	"io/ioutil"
	"encoding/json"
)

func main() {
	var s = test.CreateSquare(11)
	s.Color = "green"
	fmt.Println(s.Area())
	fmt.Println(s)

	var book = model.AuthorDto{
		UUID:      "1234",
		FirstName: "Ion",
		LastName:  "Creanga",
		Birthday:  "1850",
		Death:     "1914",
	}

	fmt.Println(book)


	booksFileContent, err := ioutil.ReadFile("model/books.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("here")
	fmt.Println(string(booksFileContent))

	err = json.Unmarshal(booksFileContent, &model.Books)
	if err != nil {
		panic(err)
	}

	fmt.Println("The deserialized books are: ")
	fmt.Println(model.Books)

	serializedBooksData, err := json.Marshal(model.Books)
	if err != nil{
		panic(err)
	}

	fmt.Println("The serialized books are:")
	fmt.Println(string(serializedBooksData))

	//////////////////////////////////////
	authorsFileContent, err := ioutil.ReadFile("model/authors.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(authorsFileContent))

	err = json.Unmarshal(authorsFileContent, &model.Authors)
	if err != nil {
		panic(err)
	}

	fmt.Println("The deserialized Authors are: ")
	fmt.Println(model.Authors)

	serializedAuthorsData, err := json.Marshal(model.Authors)
	if err != nil{
		panic(err)
	}

	fmt.Println("The serialized Authors are:")
	fmt.Println(string(serializedAuthorsData))
}
