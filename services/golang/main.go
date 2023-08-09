package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

	// Routes
	rtr := mux.NewRouter()
	// This route is added for k8s liveness probe
	rtr.HandleFunc("/", hello)
	rtr.HandleFunc("/ping", hello)
	rtr.HandleFunc("/headers", headers)
	http.Handle("/", rtr)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Println("Listening...")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Println(err)
		os.Exit(255)
	}
}
