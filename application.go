package main

import (
	"net/http"
	"time"

	"github.com/goombaio/namegenerator"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":5000", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UTC().UnixNano()
    	nameGenerator := namegenerator.NewNameGenerator(seed)

    	name := nameGenerator.Generate()
	w.Write([]byte("Hello " + name))	
}
