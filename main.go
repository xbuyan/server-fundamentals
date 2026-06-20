package main

import (
	"fmt"
	"log"
	"net/http"
	"server-fundamentals/middleware"
)

// This is a handler — a function that runs when a request arrives
// Every handler receives two things:
// w = the writer, how you send data BACK to the client
// r = the request, everything the client sent TO you
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// we explicitly check the path
	//r.URL.Path is the exact path requested by the client

	if r.URL.Path != "/" {

		//http.NotFound sends a proper 404 response
		//correct way to reject unknown routes
		http.NotFound(w, r)
		return // stops the function immediately else the code below would run
	}

	// set the staus code explicitly before writing the body
	//Headers must be set before writing the body
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Server is alive")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is a GO server built by Lucciano")
}
func main() {

	//Tell Go: When someone requests "/", run homeHandler

	http.HandleFunc("/", middleware.LoggingMiddleware(middleware.MethodMiddleware(homeHandler, "GET")))
	http.HandleFunc("/about", middleware.LoggingMiddleware(middleware.MethodMiddleware(aboutHandler, "GET")))
	// Start listening on port 9090
	// This line BLOCKS — meaning the program stays running, waiting

	log.Println("Server starting on port 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
