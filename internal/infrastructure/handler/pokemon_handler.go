package handler

import (
	"amaris/internal/app"
	"amaris/internal/infrastructure/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PokemonHandler struct {
	Service *app.PokemonService
}

func NewPokemonHandler(service *app.PokemonService) *PokemonHandler {
	return &PokemonHandler{Service: service}
}

func (h *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	log.Printf("Received request for Pokemon: %s", name)

	// Check if the Pokemon is already in the database
	existingPokemon, err := h.Service.FindPokemonByName(name)
	if err == nil {
		// Pokemon found in the database
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Pokemon %s is already in the database", existingPokemon.Name)))
		return
	} else if err != sql.ErrNoRows {
		// An error occurred while querying the database
		log.Printf("Error checking database: %v", err)
		http.Error(w, fmt.Sprintf("Error checking database: %v", err), http.StatusInternalServerError)
		return
	}

	// Pokemon not found in the database, fetch from external API
	pokemon, err := api.FetchPokemonFromAPI(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
	if err != nil {
		log.Printf("Error fetching Pokemon: %v", err)
		http.Error(w, fmt.Sprintf("Error fetching Pokemon: %v", err), http.StatusInternalServerError)
		return
	}

	// Save the Pokemon to the database
	if err := h.Service.AddPokemon(pokemon); err != nil {
		log.Printf("Error saving Pokemon: %v", err)
		http.Error(w, fmt.Sprintf("Error saving Pokemon: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Pokemon %s added successfully!", pokemon.Name)))
}
