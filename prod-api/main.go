package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cryptus-neoxys/go-micro/prod-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "prod-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/hello", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting on http://localhost:8080")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-sigChan
	l.Println("Received Terminate, Shutting down gracefully. Sig: ", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(tc); err != nil {
		l.Println("shutdown error")
	}
}
