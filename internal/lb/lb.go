package lb

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

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
	lb.controller.SetupServers(urls...)

}

func (lb *LoadBalancer) Listen(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Started listening on %s\n", addr)
	if err := http.ListenAndServe(addr, lb.controller.HTTPHandler()); err != nil {
		log.Fatal(err)
	}
}

func (lb *LoadBalancer) HealthCheck(d time.Duration) {
	log.Println("TIme after HealthCheck: ", d)
	t := time.NewTicker(d)
	for range t.C {
		log.Println("Health check starting...")
		lb.controller.HealthCheck()
		log.Println("Health check completed")
	}

}
