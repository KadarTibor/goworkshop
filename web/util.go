package web

import (
	"net/http"
	"fmt"
	"encoding/json"
)


//this function writes data to the response writer
func WriteJson(w http.ResponseWriter, data interface{}){
	serializedBooksData, err := json.Marshal(data)
	if err != nil{
		fmt.Fprintln(w,"Failed to marshal JSON data %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(serializedBooksData))
}