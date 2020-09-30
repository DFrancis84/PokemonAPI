package types

type Pokemon struct {
	Name      string    `json:"name"`
	ID        int       `json:"id"`
	Abilities []Ability `json:"abilities"`
	//Moves     []Move    `json:"moves"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type Ability struct {
	AbilityName AbilityName `json:"ability"`
	Slot        int         `json:"slot"`
}

type AbilityName struct {
	Name string `json:"name"`
}

type Move struct {
	MoveName MoveName `json:"move"`
}

type MoveName struct {
	Name string `json:"name"`
}

type Stats struct {
	AllHeights   []float64 `json:"allHeights,omitempty"`
	AllWeights   []float64 `json:"allWeights,omitempty"`
	MeanHeight   float64   `json:"meanHeight,omitempty"`
	MeanWeight   float64   `json:"meanWeight,omitempty"`
	MedianHeight float64   `json:"medianHeight,omitempty"`
	MedianWeight float64   `json:"medianWeight,omitempty"`
	ModeHeight   []float64 `json:"modeHeight,omitempty"`
	ModeWeight   []float64 `json:"modeWeight,omitempty"`
}
