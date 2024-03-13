package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	HashMap "github.com/rodblg/TinURL-Go/pkg/hash"
)

type Url struct {
}

func NewUrlController() *Url {
	return &Url{}
}

func (c *Url) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/tinyUrl", c.GetTinyUrl)
	r.Post("/", c.NumberClicks)

	return r
}

func (c *Url) GetTinyUrl(w http.ResponseWriter, r *http.Request) {

	a := HashMap.NewMap()
	log.Print(a)

	//GET ORIGINAL URL AS STRING

	//CREATE HASH TABLE BUT YOU MUST AVOID COLLISION OF THE HASH MAP

	//CALL HASH FUNCTION

	//GENERATE NEW URL
}

func (c *Url) NumberClicks(w http.ResponseWriter, r *http.Request) {

}
