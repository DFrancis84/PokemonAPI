package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DFrancis84/pokemonAPI/internal/pokemon"
	"github.com/DFrancis84/pokemonAPI/internal/types"
	"github.com/gorilla/mux"
)

type RESTAPI struct {
	Pokemon *pokemon.API
}

type Request struct {
	Pokemon []string `json:"pokemon"`
}

func New(pokemon *pokemon.API) *RESTAPI {
	return &RESTAPI{
		Pokemon: pokemon,
	}
}

func (api *RESTAPI) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (api *RESTAPI) getPokemonBios(w http.ResponseWriter, r *http.Request) {
	var req Request
	var pokemon []types.Pokemon

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	pokemon = api.Pokemon.GetBios(req.Pokemon)
	biosJSON, err := json.Marshal(pokemon)
	if err != nil {
		fmt.Printf("Error marshalling: %v\n", err)
	}
	w.Write(biosJSON)
}

func (api *RESTAPI) getPokemonStats(w http.ResponseWriter, r *http.Request) {
	var req Request
	var pokemonStats types.Stats
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	pokemonStats = api.Pokemon.GetStats(req.Pokemon)
	statsJSON, err := json.Marshal(pokemonStats)
	if err != nil {
		fmt.Printf("Error marshalling: %v\n", err)
	}
	w.Write(statsJSON)
}

func (api *RESTAPI) HandleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", api.health)
	r.HandleFunc("/pokemon/stats", api.getPokemonStats)
	r.HandleFunc("/pokemon/bios", api.getPokemonBios)

	http.ListenAndServe(":8080", r)
}
