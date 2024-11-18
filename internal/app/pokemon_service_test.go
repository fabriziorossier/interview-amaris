package app

import (
	"amaris/internal/domain"
	"testing"
)

func TestAddPokemon(t *testing.T) {
	repo := NewMockPokemonRepository()
	service := NewPokemonService(repo)

	pokemon := domain.Pokemon{
		Id:             1,
		Name:           "Bulbasaur",
		Pokemon:        "bulbasaur",
		Height:         7,
		Weight:         69,
		BaseExperience: 64,
	}

	err := service.AddPokemon(pokemon)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	savedPokemon, err := repo.FindByName(pokemon.Name)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if savedPokemon.Name != "Bulbasaur" {
		t.Fatalf("expected Bulbasaur, got %v", savedPokemon.Name)
	}
}

func TestFindPokemonByName(t *testing.T) {
	repo := NewMockPokemonRepository()
	service := NewPokemonService(repo)

	pokemon := domain.Pokemon{
		Id:             1,
		Name:           "Bulbasaur",
		Pokemon:        "bulbasaur",
		Height:         7,
		Weight:         69,
		BaseExperience: 64,
	}

	service.AddPokemon(pokemon)

	foundPokemonByName, err := service.FindPokemonByName(pokemon.Name)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if foundPokemonByName.Name != "Bulbasaur" {
		t.Fatalf("expected Bulbasaur, got %v", foundPokemonByName.Name)
	}
}
