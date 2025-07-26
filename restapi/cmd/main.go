package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/schnell18/play-golang/restapi/internal/config"
	"github.com/schnell18/play-golang/restapi/internal/route"
)

func main() {
	// load app config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}

	// set up the http server
	mux := http.NewServeMux()

	// setup routes
	route.SetupRoutes(mux)

	// server instance
	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	fmt.Printf("Serving at %s\n", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed %v", err)
	}
}
