package handlers

import (
	"amaris/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var baseURL = "https://pokeapi.co/api/v2/pokemon/"

func PokemonHandler(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/")
	if param == "" {
		http.Error(w, "Missing parameter in URL path", http.StatusBadRequest)
		return
	}

	fullURL := baseURL + param

	resp, err := http.Get(fullURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	var pokemon models.Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing JSON: %v", err), http.StatusInternalServerError)
		return
	}

	var response string
	if pokemon.Name == param || pokemon.Pokemon == param {
		caser := cases.Title(language.English)
		createdPokemon := caser.String(param)
		response = fmt.Sprintf("<html><body><h1>Pokemon created: %s</h1></body></html>", createdPokemon)
	} else {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}
