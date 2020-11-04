package main

import (
	"context"
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"os"
	"os/signal"
)

// newServer sets up a localhost server using the gorilla/mux package
// and calls handlers for endpoints.
func newServer(port string) *myServer {

	s := &myServer{
		Server: http.Server{
			Addr:         (":" + port),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		shutdownReq: make(chan bool),
	}

	router := mux.NewRouter()

	// Register handlers
	router.HandleFunc("/api", topLevel)
	router.HandleFunc("/api/{Id=threads}", getThreads)
	router.HandleFunc("/shutdown", s.shutdownHandler)

	// Set http server handler
	s.Handler = router

	return s
}

// waitShutdown will perform a server shutdown either through API call or interrupt.
func (s *myServer) waitShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	// Wait interrupt or shutdown request through /shutdown.
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}

	log.Printf("Stoping http server ...")

	// Create shutdown context with 10 second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server.
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
}
