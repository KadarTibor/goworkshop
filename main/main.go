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


	fileContent, err := ioutil.ReadFile("model/books.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("here")
	fmt.Println(string(fileContent))

	var books []model.BookDto
	err = json.Unmarshal(fileContent, &books)
	if err != nil {
		panic(err)
	}

	fmt.Println("The deserialized books are: ")
	fmt.Println(books)

	serializedData, err := json.Marshal(books)
	if err != nil{
		panic(err)
	}

	fmt.Println("The serialized books are:")
	fmt.Println(string(serializedData))
}
