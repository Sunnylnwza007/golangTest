package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "log"
	
)

type Test struct {
	Events []Event
	Destination string
}
type Event struct{
	Type string
	ReplyToken string
	Source Source
	Timestamp int64
	Mode string
	Message Message

}

type Source struct{
	UserId string
	Type string
	Address string
	Latitude float32
	Longitude float32
}


type Message struct{
	Type string
	Id string
	
}

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
	
	// addBook := test{
	// 			 Events: "test",
	// 			 Destination: "1234",
	// 		   }
	// json.NewEncoder(w).Encode(addBook) // (1)
	// decoder := json.NewDecoder(r.Body)
	// var t test
	// err := decoder.Decode(&t)
	// if err != nil {
    //     panic(err)
    // }
	// log.Println(t)
	var p Test

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the Person struct...
    fmt.Fprintf(w, "Person: %+v", p.Events)
	
 }

 func main() {
	handReq() 
}