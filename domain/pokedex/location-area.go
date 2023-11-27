package pokedex

type LocationArea struct {
	Name string `json:"name"`
}

type NextLocationAreasFunc = func() error
type PreviousLocationAreasFunc = func() error
type GetLocationAreasFunc = func() []LocationArea
