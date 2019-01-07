package main

import (
	"log"
	"net/http"

	"github.com/javipolo/telehooker/handlers"
)

func main() {
	http.HandleFunc("/healthz", handlers.Healthz)
	http.HandleFunc("/wormly", handlers.Wormly)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
