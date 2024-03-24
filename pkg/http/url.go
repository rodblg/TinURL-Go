package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Url struct {
}

func NewUrlController() *Url {
	return &Url{}
}

func (c *Url) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/tinyUrl", c.GetTinyUrl)
	//r.Post("/", c.NumberClicks)

	return r
}

func (c *Url) GetTinyUrl(w http.ResponseWriter, r *http.Request) {

	//a := NewMap()

	//log.Print(a)

	//GET ORIGINAL URL AS STRING

	//CREATE HASH TABLE BUT YOU MUST AVOID COLLISION OF THE HASH MAP

	//CALL HASH FUNCTION

	//GENERATE NEW URL
}

func (c *Url) GenerateTiny(w http.ResponseWriter, r *http.Request) {

	//get url from request context

	originalUrl := r.Header.Get("url")
	if originalUrl == "" {
		//Return err response for empty url to hash
		return
	}

	//we call our hash function to transform the url
	tinyUrl := HashURL(originalUrl)

	//we call our usecase for writing into the database
	responseJson, err := usecases.NewUrl(tinyUrl)
	if err != nil {
		//return err response
		return
	}
	//we return json with key, original and new url
	render.Status(r, http.StatusOK)
	render.JSON(responseJson)
}
