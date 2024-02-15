package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	items  = []Person{}
	NextID = 1
)

func GoHandler(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "John", Age: 25, Id: 1}

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

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.Id = NextID
	NextID++

	ps := append(items, p)
	err = json.NewEncoder(w).Encode(ps)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", GoHandler)
	mux.HandleFunc("POST /add", AddHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
