package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListenAndServe() {

	log.Println("running server...")
	r := chi.NewRouter()

	r.Mount("/urls", NewUrlController().Routes())

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println("error initializing server...")
		return
	}
}
