package web

import (
"net/http"
"os"
"fmt"
"goworkshop/model"
	"encoding/json"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	http.HandleFunc("/books", httpHandler)
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

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}