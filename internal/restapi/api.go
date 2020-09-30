package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DFrancis84/pokemonAPI/internal/pokemon"
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

func (api *RESTAPI) getPokemonStats(w http.ResponseWriter, r *http.Request) {
	var pokemonStat *pokemon.Stat

	pokemon := r.URL.Query().Get("name")
	pokemonStat = api.Pokemon.GetStat(pokemon)
	statsJSON, err := json.Marshal(pokemonStat)
	if err != nil {
		fmt.Printf("Error marshalling: %v\n", err)
	}
	w.Write(statsJSON)
}

func (api *RESTAPI) getMultiplePokemonStats(w http.ResponseWriter, r *http.Request) {
	var req Request
	var pokemonStats *pokemon.Stats
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
	r.HandleFunc("/pokemon", api.getPokemonStats)
	r.HandleFunc("/list/pokemon", api.getMultiplePokemonStats)

	http.ListenAndServe(":8080", r)
}
