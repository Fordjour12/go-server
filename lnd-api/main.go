package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func GoHandler(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "John", Age: 25}

	// Encode the struct to JSON and write it to the response(w.WriteHeader(http.StatusOK)is not needed)
	// err := json.NewEncoder(w).Encode(p)

	// NOTE: This is the same as using json.NewEncoder(w).Encode(p)
	jsonData, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", GoHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
