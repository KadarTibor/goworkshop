package web

import (
"net/http"
"os"
"fmt"
"goworkshop/model"
	"encoding/json"
	"strings"
	"github.com/gorilla/mux"
	"io/ioutil"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	gorillaMux := mux.NewRouter()
	http.HandleFunc("/books", httpHandler)
	http.HandleFunc("/authors/",getAuthorByUUID)
	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		serializedBooksData, err := json.Marshal(model.Books)
		if err != nil{
			panic(err)
		}
		fmt.Fprintln(w, string(serializedBooksData))
		break

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Not supported method!")
	}

}

func getAuthorByUUID(w http.ResponseWriter, r *http.Request){
	path := r.URL.Path  //author/UUId

	uuid :=	strings.TrimPrefix(path,"/authors/")

	found := false
	for _, author := range model.Authors {
		if author.UUID == uuid{
			serializedBooksData, err := json.Marshal(author)
			if err != nil{
				panic(err)
			}
			fmt.Fprintln(w, string(serializedBooksData))
			found = true
			}
	}

	w.WriteHeader(http.StatusNotFound)
	if(!found){
		fmt.Fprintln(w,"The author with the UUID %s does not exist", uuid)
	}

}

func CreateAuthor(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if(err!= nil){
		fmt.Fprintf(w, "{\"errorMessage\":\"%s\"}", err.Error())
		return
	} else {
			//TODO
	}
}


//TODO
//implement all the methods for author and book
func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}