package main

import (
	"log"
	"net/http"

	"url-shortner/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /encode", handler.Encode)
	mux.HandleFunc("POST /decode", handler.Decode)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
