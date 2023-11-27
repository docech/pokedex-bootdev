package print

import (
	"fmt"

	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

func PrintLocationAreas(locationAreas []pokedex.LocationArea) {
	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}
}