package app

import (
	"amaris/internal/domain"
	"database/sql"
	"strings"
)

type MockPokemonRepository struct {
	Pokemons map[int]domain.Pokemon
}

func NewMockPokemonRepository() *MockPokemonRepository {
	return &MockPokemonRepository{Pokemons: make(map[int]domain.Pokemon)}
}

func (m *MockPokemonRepository) Save(pokemon domain.Pokemon) error {
	m.Pokemons[pokemon.Id] = pokemon
	return nil
}

func (m *MockPokemonRepository) FindByName(name string) (domain.Pokemon, error) {
	for _, pokemon := range m.Pokemons {
		if strings.EqualFold(pokemon.Name, name) {
			return pokemon, nil
		}
	}
	return domain.Pokemon{}, sql.ErrNoRows
}
