package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/ping", hello)
	http.HandleFunc("/headers", headers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Println("Listening...")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
