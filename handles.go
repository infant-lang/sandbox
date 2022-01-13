package main

import (
	"encoding/json"
	"net/http"
)

type Code struct {
	Code []string `json:"code"`
}

type Output struct {
	Output string `json:"output"`
}

func StartSandbox(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ready"))
}

func PlayHandler(w http.ResponseWriter, r *http.Request) {

	defer func() {     
		if e := recover(); e != nil {
			
			var output string = message + "\n" + panicMessage

			outputJson, _ := json.Marshal(Output{output})
			w.Header().Set("Content-Type", "application/json")
			w.Write(outputJson)

		} else {

			var output string = message

			outputJson, _ := json.Marshal(Output{output})
			w.Header().Set("Content-Type", "application/json")
			w.Write(outputJson)
		}

		message = ""
		panicMessage = ""
	}()

	var code Code
	_ = json.NewDecoder(r.Body).Decode(&code)

	p := 0
	m := 0

	for lineNumber, eachline := range code.Code {
		p, m = parse(eachline, lineNumber + 1, p, m)
	}


}