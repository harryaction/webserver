package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/weather/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "Response")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
