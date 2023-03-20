package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./example")))
	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		log.Fatal(err)
	}
	c := make(chan bool)
	<-c

}
