package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListenAndServe() {
	r := chi.NewRouter()

	r.Mount("/urls", NewUrlController().Routes())

	http.ListenAndServe(":8080", r)
}
