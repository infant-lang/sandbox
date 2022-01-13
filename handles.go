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

	p := 0
	m := 0

	for lineNumber, eachline := range code.Code {
		p, m = parse(eachline, lineNumber + 1, p, m)
	}


}