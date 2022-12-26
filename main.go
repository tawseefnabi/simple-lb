package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/simple-lb/internal/lb"
	"github.com/simple-lb/internal/utils"
	log "github.com/sirupsen/logrus"
)

var (
	flagURL utils.FlagURL
	port    int
)

func init() {
	flag.Var(&flagURL, "servers", "Use commas to separate")
	flag.IntVar(&port, "port", 8090, "Port to serve")
	flag.Parse()
	if len(flagURL.URLs) == 0 {
		log.Fatal("Please provide one or more servers to load balance")
	}
}
func main() {
	fmt.Println("Simple lb")
	l := lb.New()
	l.Register(flagURL.URLs...)
	go l.HealthCheck(1 * time.Second)
	l.Listen(port)

}
