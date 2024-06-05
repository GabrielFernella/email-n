package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		obj := map[string]string{"message": "success"}

		render.JSON(w, r, obj)
	})

	http.ListenAndServe(":3000", r)
}