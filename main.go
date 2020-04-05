package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type ResponseJson struct {
	//provides a struct to response with pseudo-random values
	Id      int    `json:"id"` //mentioned as best practice to cast struct into json
	Feeling string `json:"feeling"`
}

func randomInt() int {
	//returns random int in range 0-500
	return rand.Intn(500)
}

func randomString() string {
	//returns random string 20 symbol length
	bytes := make([]byte, 20)
	for i := 0; i < 20; i++ {
		bytes[i] = byte(randomInt())
	}
	return string(bytes)
}

func main() {
	//not sure 'bout doc string here, but
	//it's a basic web server implementation, which is listen on 8080 and should response a JSON
	//provided by Response func
	rand.Seed(time.Now().UnixNano()) //Seeding seeds for random
	http.HandleFunc("/weather/", func(w http.ResponseWriter, request *http.Request) {
		response := ResponseJson{Id: randomInt(), Feeling: randomString()} //generating response
		json.NewEncoder(w).Encode(response)                                //write encoded response to response writer

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
