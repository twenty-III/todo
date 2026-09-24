package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Envelope map[string]any

func ReadJson[T any](r *http.Request) (T, error) {
	var res T
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}

func WriteJson[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Print(err)
	}
}
