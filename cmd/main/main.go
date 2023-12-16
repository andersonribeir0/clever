package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andersonribeir0/clever/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	router := routes.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("HTTP_LISTEN_ADDR")
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server listening on port %s\n", port)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
