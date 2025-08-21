package main

import (
	"log"
	"net/http"

	"github.com/schnell18/play-golang/ogen/ping"
)

func main() {
	service := ping.NewPingService()
	srv, err := ping.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	// And we serve HTTP until the world ends.
	if err = http.ListenAndServe(":7000", srv); err != nil {
		log.Fatal(err)
	}
}
