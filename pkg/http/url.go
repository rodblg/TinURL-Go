package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Url struct {
}

func NewUrlController() *Url {
	return &Url{}
}

func (c *Url) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", c.GetTinyUrl)
	r.Post("/", c.NumberClicks)

	return r
}

func (c *Url) GetTinyUrl(w http.ResponseWriter, r *http.Request) {

}

func (c *Url) NumberClicks(w http.ResponseWriter, r *http.Request) {

}
