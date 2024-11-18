package handler

import (
	"amaris/internal/app"
	"amaris/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetPokemon(t *testing.T) {
	repo := app.NewMockPokemonRepository()
	service := app.NewPokemonService(repo)
	handler := NewPokemonHandler(service)

	pokemon := domain.Pokemon{
		Id:             1,
		Name:           "Bulbasaur",
		Pokemon:        "bulbasaur",
		Height:         7,
		Weight:         69,
		BaseExperience: 64,
	}

	service.AddPokemon(pokemon)

	req, err := http.NewRequest("GET", "/pokemon/bulbasaur", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/pokemon/{name}", handler.GetPokemon).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Pokemon Bulbasaur is already in the database"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
