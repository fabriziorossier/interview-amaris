package persistence

import (
	"amaris/internal/domain"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Save(pokemon domain.Pokemon) error {
	_, err := r.db.Exec("INSERT INTO pokemon (id, name, pokemon, height, weight, base_experience) VALUES ($1, $2, $3, $4, $5, $6)", pokemon.Id, pokemon.Name, pokemon.Pokemon, pokemon.Height, pokemon.Weight, pokemon.BaseExperience)
	return err
}

func (r *PostgresRepository) FindByName(name string) (domain.Pokemon, error) {
	var pokemon domain.Pokemon
	err := r.db.QueryRow("SELECT id, name, pokemon, height, weight, base_experience FROM pokemon WHERE name = $1 OR pokemon = $1", name).Scan(&pokemon.Id, &pokemon.Name, &pokemon.Pokemon, &pokemon.Height, &pokemon.Weight, &pokemon.BaseExperience)
	return pokemon, err
}
