package models

type Url struct {
	Key         string `json:"key"`
	OriginalUrl string `json:"original_url"`
	TinyUrl     string `json:"tiny_url"`
}
