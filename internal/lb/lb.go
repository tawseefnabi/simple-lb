package lb

import (
	"fmt"
	"log"
	"net/url"

	"github.com/simple-lb/internal/server"
)

type LoadBalancer struct {
	controller *server.Controller
}

func New() *LoadBalancer {
	fmt.Println("New func from lb called...")
	return &LoadBalancer{
		controller: server.NewController(),
	}
}

func (lb *LoadBalancer) Register(urls ...*url.URL) {
	fmt.Println("Regsiter.. urls: ", urls)
	for _, u := range urls {
		log.Printf("configured server: %s", u)
	}

}
