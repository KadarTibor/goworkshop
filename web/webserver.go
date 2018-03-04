package web

import (
	"net/http"
	"os"
	"fmt"
	"goworkshop/model"
	"encoding/json"
	"github.com/gorilla/mux"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	router := mux.NewRouter()
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
	var port = getPort()
	fmt.Println("+-------------------------------")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("-------------------------------+")
	if err := http.ListenAndServe(":" + port, router); err != nil {
		panic(err)
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
