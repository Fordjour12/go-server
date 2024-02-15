package main

import (
	"log"
	"net/http"
)

//func RenderHandler(w http.ResponseWriter, r *http.Request) {
//	http.ServeFile(w, r, "static/")
//}

func main() {
	srv := http.FileServer(http.Dir("static"))
	http.Handle("GET /", srv)

	//	srv := http.NewServeMux()
	//	srv.HandleFunc("GET /", RenderHandler)

	log.Fatal(http.ListenAndServe(":8080", srv))
}
