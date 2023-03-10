package server

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

type server struct {
	url     *url.URL
	handler http.Handler
}

func newServer(u *url.URL, handler http.Handler) *server {
	return &server{url: u, handler: handler}
}

func (s *server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.handler.ServeHTTP(rw, req)
}
func (s *server) IsAlive() bool {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", s.url.Host, timeout)
	return err == nil
}
