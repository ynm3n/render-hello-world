package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)
	s := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("yo")
	fmt.Fprint(w, "hello, world")
}
