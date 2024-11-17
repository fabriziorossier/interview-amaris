package api

import (
	"amaris/internal/domain"
	"encoding/json"
	"net/http"
)

func FetchPokemonFromAPI(url string) (domain.Pokemon, error) {
	resp, err := http.Get(url)
	if err != nil {
		return domain.Pokemon{}, err
	}
	defer resp.Body.Close()

	var pokemon domain.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return domain.Pokemon{}, err
	}

	return pokemon, nil
}
