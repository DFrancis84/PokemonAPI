package main

import (
	"github.com/DFrancis84/pokemonAPI/internal/pokemon"
	"github.com/DFrancis84/pokemonAPI/internal/restapi"
)

func main() {
	pokemon := pokemon.New()

	api := restapi.New(pokemon)
	api.HandleRequests()
}
