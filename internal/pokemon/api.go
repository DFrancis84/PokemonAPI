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

type Stat struct {
	Name   string
	Height int
	Weight int
}

type Stats struct {
	Pokemon      []string
	AllHeights   []float64
	AllWeights   []float64
	MeanHeight   float64
	MeanWeight   float64
	MedianHeight float64
	MedianWeight float64
	ModeHeight   []float64
	ModeWeight   []float64
}

func New() *API {
	client := &http.Client{}
	return &API{Client: client}
}

func (api *API) GetStat(name string) *Stat {
	var result types.Result

	url := fmt.Sprintf("%v/%v", apiURL, name)

	res, err := api.Client.Get(url)
	if err != nil {
		fmt.Printf("Error getting response from API: %v", err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}

	return &Stat{
		Name:   result.Name,
		Height: result.Height,
		Weight: result.Weight,
	}
}

func (api *API) GetStats(names []string) *Stats {
	var meanHeight, meanWeight, medianHeight, medianWeight float64
	var modeHeight, modeWeight []float64
	var result types.Result

	heights := []float64{}
	weights := []float64{}

	for _, x := range names {
		url := fmt.Sprintf("%v/%v", apiURL, x)

		res, err := api.Client.Get(url)
		if err != nil {
			fmt.Printf("Error getting response from API: %v", err)
		}

		body, _ := ioutil.ReadAll(res.Body)

		defer res.Body.Close()
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println(err)
		}

		heights = append(heights, float64(result.Height))
		weights = append(weights, float64(result.Weight))
	}

	meanHeight, meanWeight = getMean(heights, weights)
	medianHeight, medianWeight = getMedian(heights, weights)
	modeHeight, modeWeight = getMode(heights, weights)

	return &Stats{
		Pokemon:      names,
		AllHeights:   heights,
		AllWeights:   weights,
		MedianHeight: medianHeight,
		MedianWeight: medianWeight,
		MeanHeight:   meanHeight,
		MeanWeight:   meanWeight,
		ModeHeight:   modeHeight,
		ModeWeight:   modeWeight,
	}

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
