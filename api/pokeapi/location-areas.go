package pokeapi

import (
	"errors"

	"github.com/docech/pokedex-bootdev/api"
	"github.com/docech/pokedex-bootdev/api/http"
	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

type locationAreasResource struct {
	NextResults     any         `json:"next"`
	PreviousResults any         `json:"previous"`
	Results         []pokedex.LocationArea `json:"results"`
}

func fetchResource(url string) (api.Resource[[]pokedex.LocationArea], error) {
	var resource locationAreasResource
	err := http.Fetch[[]pokedex.LocationArea](url, &resource)
	return &resource, err
}

func NewLocationAreasResource() api.Resource[[]pokedex.LocationArea] {
	return &locationAreasResource{
		NextResults: "https://pokeapi.co/api/v2/location-area/",
		PreviousResults: nil,
		Results: []pokedex.LocationArea{},
	}
}

func (r *locationAreasResource) Next() (api.Resource[[]pokedex.LocationArea], error) {
	nextURL, ok := r.NextResults.(string)
	if !ok {
		return nil, errors.New("no next resource")
	}
	return fetchResource(nextURL)
}

func (r *locationAreasResource) Previous() (api.Resource[[]pokedex.LocationArea], error) {
	previousURL, ok := r.PreviousResults.(string)
	if !ok {
		return nil, errors.New("no previous resource")
	}
	return fetchResource(previousURL)
}

func (r *locationAreasResource) Data() []pokedex.LocationArea {
	return r.Results
}

