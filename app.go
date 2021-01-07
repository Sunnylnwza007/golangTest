package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"os"
	// "log"
	"context"
	"sunny/test-golang/models"
	"sunny/test-golang/db"
	"github.com/joho/godotenv"
	"github.com/go-redis/redis/v8"
	// "github.com/go-redis/redis/v8"
)

type Database struct {
    Redis context.Context
}

var ctx = context.Background()
func handReq(){
	rdb := db.GetRedis().Client;
	redis := context.WithValue(ctx, "redis", rdb)	

	database := &Database{Redis: redis}

	fmt.Println(database)
	//connect

	http.HandleFunc("/", database.addHandler)
	http.HandleFunc("/test", testJson)

  http.ListenAndServe(":8080", nil)
}

  
func (db *Database) addHandler(w http.ResponseWriter, r *http.Request) {
	redis := db.Redis.Value("redis").(*redis.Client)
	data, err := redis.Get(context.Background(),"key").Result()
	if err != nil {
		    panic(err)
		}
	fmt.Fprint(w, data)
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	test := os.Getenv("SAMPLE_ENV")
	fmt.Println(test)

	
	handReq() 
}