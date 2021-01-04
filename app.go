package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "log"
	"sunny/test-golang/models"
	"github.com/fatih/color"
)

func handReq(){
	

	//connect

	http.HandleFunc("/", addHandler)
	// http.HandleFunc("/test", addHandler)
	http.HandleFunc("/test", testJson)
	http.ListenAndServe(":8080", nil)
	
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test success")
	fmt.Println("Hello World")
}

func testJson(w http.ResponseWriter, r *http.Request) {
	var p models.Test

    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the Person struct...
    fmt.Fprintf(w, "Person: %+v", p.Events[0].Message)
	
 }

 func main() {
	color.Blue("%s \n", "Hello module 1")
	handReq() 
}