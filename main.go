package main

import (
	"log"
	"net/http"
)

func main() {

	log.Println("Sentra starting on port 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
