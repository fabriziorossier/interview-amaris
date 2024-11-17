package app

import "amaris/internal/domain"

type PokemonService struct {
	repo PokemonRepository
}

func NewPokemonService(repo PokemonRepository) *PokemonService {
	return &PokemonService{repo: repo}
}

func (s *PokemonService) FindPokemonByName(name string) (domain.Pokemon, error) {
	return s.repo.FindByName(name)
}

func (s *PokemonService) AddPokemon(pokemon domain.Pokemon) error {
	return s.repo.Save(pokemon)
}
