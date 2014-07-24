package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	log.Println("Starting Server on", port)
	http.HandleFunc("/", all)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func all(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello GopherCasts!"))
}
