package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/docech/pokedex-bootdev/api"
)

func Fetch[V any](url string, resource api.Resource[V]) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}

	return nil
}