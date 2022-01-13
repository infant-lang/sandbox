package main

import (
	"encoding/json"
	"net/http"
)

func StartSandbox(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ready"))
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {

	// get the json object from the request
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	// return the data as a json object
	json.NewEncoder(w).Encode(data)

}