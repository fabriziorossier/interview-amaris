package main

import (
	"amaris/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.PokemonHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
