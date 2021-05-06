package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have visited %s!", r.URL.Path)
}

func main() {
	port := "8080"
	http.HandleFunc("/", handle)
	fmt.Println("Starting web server on port " + port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
