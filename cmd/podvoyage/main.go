package main

import (
	"log"
	"net/http"

	"github.com/airelcamilo/podvoyage-backend/internal/pkg/router"
	"github.com/rs/cors"
)

func main() {
	// Listening to port
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "DELETE"},
	})
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", c.Handler(r)))

	// go build .
	// go run main.go
}
