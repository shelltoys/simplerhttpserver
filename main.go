package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	var port string

	if len(os.Args) < 2 {
		port = "8000"
	} else {
		port = os.Args[1]
	}

	log.Println("Now listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
