package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("New REST APIs in Go!")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Return all comments")

	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		fmt.Fprintf(w, "Return a single comment id: %s", id)

	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Post a new comment")

	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}

}
