package lb

import (
	"fmt"
	"github.com/simple-lb/internal/server"
	"log"
	"net/http"
	"net/url"
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
	lb.controller.SetupServers(urls...)

}

func (lb *LoadBalancer) Listen(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Started listening on %s\n", addr)
	if err := http.ListenAndServe(addr, lb.controller.HTTPHandler()); err != nil {
		log.Fatal(err)
	}
}
