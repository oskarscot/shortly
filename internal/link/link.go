package link

import "errors"

var (
	ErrLinkNotFound = errors.New("link not found")
	ErrInvalidUrl   = errors.New("invalid url")
)

type LinkCreateRequest struct {
	Url string `json:"url"`
}

type Link struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Url  string `json:"url"`
}
