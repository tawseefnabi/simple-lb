package server

import (
	"container/list"
	"net/url"
)

type Controller struct {
	servers map[int]*server
	upIDs   *list.List
	dowsIDs *list.List
	// mux     sync.Mutex
}

func NewController() *Controller {
	return &Controller{
		upIDs:   list.New(),
		dowsIDs: list.New(),
	}
}

func (c *Controller) SetUpServers(urls ...*url.URL) {
	c.servers = make(map[int]*server, len(urls))
	for _, u := range urls {
		id := i + 1
		c.servers[id] = newServer
	}
}
