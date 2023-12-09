package main

import (
	"fmt"
	"net/http"

	"github.com/andersonribeir0/clever/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on port %d\n", port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
