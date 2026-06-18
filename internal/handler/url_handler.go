package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Encode func(http.ResponseWriter, *http.Request)
	Decode func(http.ResponseWriter, *http.Request)
}

func Encode(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"short_url": "http://localhost:8080/abc",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Decode(w http.ResponseWriter, r *http.Request) {

}
