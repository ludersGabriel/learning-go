package main

import (
	"ex/server/internal/routes"
	"fmt"
	"net/http"
)

func main() {
	router := routes.Router()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on %s\n", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
