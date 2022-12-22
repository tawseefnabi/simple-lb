package server

import (
	"net/http"
	"net/url"
)

type server struct {
	url     *url.URL
	handler *http.Handler
}

func newServer(u *url.URL, handler *http.Handler) *server {
	return &server{url: u, handler: handler}
}
