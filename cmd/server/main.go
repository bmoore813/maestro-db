package main

import (
	"log"

	"github.com/bmoore813/maestro-db/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
