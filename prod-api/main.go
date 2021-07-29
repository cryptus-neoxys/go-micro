package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cryptus-neoxys/go-micro/prod-api/handlers"
)


func main() {
	l := log.New(os.Stdout, "prod-api", log.LstdFlags)

	hh := handlers.NewHello(l)
    gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

    http.ListenAndServe(":9091", sm)
}