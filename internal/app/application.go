package app

import "amaris/internal/domain"

type PokemonRepository interface {
	Save(pokemon domain.Pokemon) error
	FindByName(name string) (domain.Pokemon, error)
}
