package main

import (
	"encoding/json"
	"net/http"
)

var VERSION = "0.0.1"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]string{
			"version": VERSION,
		}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			return
		}
	})
	s := http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
