package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	if err := http.ListenAndServe(":8080", chi.NewRouter()); err != nil {
		panic(err)
	}
}
