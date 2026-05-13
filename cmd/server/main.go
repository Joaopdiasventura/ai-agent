package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Joaopdiasventura/ai-agent/server"
)

func main() {
	handler, err := server.Handler()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("server running on http://localhost:" + port)

	err = http.ListenAndServe(":"+port, handler)

	if err != nil {
		log.Fatal(err)
	}
}
