package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var message string =  ""
var panicMessage string = ""


func main() {

	var port string = GoDotEnvVariable("PORT")

	// Init the Mux router
	r := mux.NewRouter()


	// Route handlers / Endpoints
	r.HandleFunc("/api/1.0.1/play", StartSandbox).Methods("GET")
	r.HandleFunc("/api/1.0.1/play", PlayHandler).Methods("POST")

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000", "http://localhost:3003", "https://infant-lang-playground.netlify.app"},
        AllowCredentials: true,
    })

	// Adding CORS to the router
	handler := c.Handler(r)
	

	// Start the server
	fmt.Println("Server listening on port "+ port)
	log.Fatal(http.ListenAndServe(":" + port, handler))
}

