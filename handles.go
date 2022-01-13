package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Code struct {
	Code []string `json:"code"`
}

func StartSandbox(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ready"))
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	var code Code
	_ = json.NewDecoder(r.Body).Decode(&code)

	// loop through the code array and print it
	for _, line := range code.Code {
		fmt.Println(line)
	}

}