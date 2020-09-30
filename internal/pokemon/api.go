package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/montanaflynn/stats"

	"github.com/DFrancis84/pokemonAPI/internal/types"
)

var (
	apiURL = "https://pokeapi.co/api/v2/pokemon"
)

type API struct {
	Client *http.Client
}

func New() *API {
	client := &http.Client{}
	return &API{Client: client}
}

func (api *API) GetBios(names []string) []types.Pokemon {
	bios := []types.Pokemon{}

	for _, x := range names {
		var pokemon types.Pokemon
		url := fmt.Sprintf("%v/%v", apiURL, x)

		res, err := api.Client.Get(url)
		if err != nil {
			fmt.Printf("Error getting response from API: %v", err)
		}

		body, _ := ioutil.ReadAll(res.Body)

		defer res.Body.Close()
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			fmt.Println(err)
		}
		bios = append(bios, pokemon)
	}

	return bios
}

func (api *API) GetStats(names []string) types.Stats {
	stats := types.Stats{}

	for _, x := range names {
		var pokemon types.Pokemon
		url := fmt.Sprintf("%v/%v", apiURL, x)

		res, err := api.Client.Get(url)
		if err != nil {
			fmt.Printf("Error getting response from API: %v", err)
		}
		body, _ := ioutil.ReadAll(res.Body)

		defer res.Body.Close()
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			fmt.Println(err)
		}
		stats.AllHeights = append(stats.AllHeights, float64(pokemon.Height))
		stats.AllWeights = append(stats.AllWeights, float64(pokemon.Weight))
	}

	stats.MeanHeight, stats.MeanWeight = getMean(stats.AllHeights, stats.AllWeights)
	stats.MedianHeight, stats.MedianWeight = getMedian(stats.AllHeights, stats.AllWeights)
	stats.ModeHeight, stats.ModeWeight = getMode(stats.AllHeights, stats.AllWeights)

	return stats
}

func getMedian(heights, weights []float64) (float64, float64) {
	var medianHeight, medianWeight float64
	medianHeight, _ = stats.Median(heights)
	medianWeight, _ = stats.Median(weights)
	return medianHeight, medianWeight
}

func getMean(heights, weights []float64) (float64, float64) {
	var meanHeight, meanWeight float64
	meanHeight, _ = stats.Mean(heights)
	meanWeight, _ = stats.Mean(weights)
	return meanHeight, meanWeight
}

func getMode(heights, weights []float64) ([]float64, []float64) {
	var modeHeight, modeWeight []float64
	modeHeight, _ = stats.Mode(heights)
	modeWeight, _ = stats.Mode(weights)
	return modeHeight, modeWeight
}
