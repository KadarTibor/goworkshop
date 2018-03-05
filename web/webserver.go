package web

import (
	"net/http"
	"os"
	"fmt"
	"github.com/gorilla/mux"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	router := mux.NewRouter()
	for _, route := range routes {
		handleFunc := log(route.HandlerFunc)
		router.HandleFunc(route.Pattern, handleFunc).Methods(route.Method)
	}
	var port = getPort()
	fmt.Println("+-------------------------------")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("-------------------------------+")
	if err := http.ListenAndServe(":" + port, router); err != nil {
		panic(err)
	}
}

func log(funcHandler http.HandlerFunc) http.HandlerFunc{
	fmt.Println("Returning the function")
	return func (rw http.ResponseWriter, r *http.Request){
		fmt.Println("New REST request " + r.URL.Path)
		funcHandler(rw, r)
		fmt.Println("REST request ended")
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
