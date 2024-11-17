package main

import (
	"amaris/internal/app"
	"amaris/internal/infrastructure/handler"
	"amaris/internal/infrastructure/persistence"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		getEnv("PGUSER", "postgres"),
		getEnv("PGPASSWORD", "postgres"),
		getEnv("PGDATABASE", "postgres"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	repo := persistence.NewPostgresRepository(db)
	service := app.NewPokemonService(repo)
	pokemonHandler := handler.NewPokemonHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{name}", pokemonHandler.GetPokemon).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
