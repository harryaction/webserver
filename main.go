package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" //const provides charset for random string func

type DataJson struct {
	//provides a struct for data in the weather object
	Id      int    `json:"id"` //mentioned as best practice to cast struct into json
	Feeling string `json:"feeling"`
}

type ResponseJson struct {
	//provides a struct for weather object in response
	Weather DataJson `json:"weather"`
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano())) //just to randomize our string

func randomInt() int {
	//returns random int in range 0-500
	return rand.Intn(500)
}

func String(charset string) string {
	b := make([]byte, 20) //creating a slice of 20 (yeah, dynamic arrays, bro)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))] //getting random symbol from charset
	}
	return string(b)
}
func Server(w http.ResponseWriter, request *http.Request) {
	response := ResponseJson{Weather: DataJson{Id: randomInt(), Feeling: String(charset)}} //generating response
	json.NewEncoder(w).Encode(response)                                                    //write encoded response to response writer
}
func main() {
	//not sure 'bout doc string here, but
	//it's a basic web server implementation, which is listen on 8080 and should response a JSON
	//provided by Response func
	http.HandleFunc("/weather/", Server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
