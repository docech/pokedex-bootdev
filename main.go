package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/docech/pokedex-bootdev/api/http"
	"github.com/docech/pokedex-bootdev/api/pokeapi"
	"github.com/docech/pokedex-bootdev/cli/commands"
	"github.com/docech/pokedex-bootdev/domain/pokedex"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := commands.NewCliCommands()
	locationAreaResource := pokeapi.NewLocationAreasResource(
		"https://pokeapi.co/api/v2/location-area/", 
		http.CacheConfig {
			MaxAge: 60 * time.Second,
		},
	)

	cmds.Register(commands.NewHelpCommand(commands.HelpDeps{
		ProvideAbouts: cmds.About,
	}))
	cmds.Register(commands.NewMapCommand(commands.MapResources{
		ProvideLocationAreas: func () ([]pokedex.LocationArea, error) {
			nextResource, err := locationAreaResource.Next()
			
			if err != nil {
				return nil, err
			}

			locationAreaResource = nextResource

			return nextResource.Data(), nil
		},
	}))
	cmds.Register(commands.NewMapbCommand(commands.MapResources{
		ProvideLocationAreas: func () ([]pokedex.LocationArea, error) {
			prevResource, err := locationAreaResource.Previous()
			
			if err != nil {
				return nil, err
			}

			locationAreaResource = prevResource

			return prevResource.Data(), nil
		},
	}))
	cmds.Register(commands.NewExitCommand())
	
	fmt.Println("Starting Pokedex...")
	for {
		fmt.Print("pokedex > ")
		
		scanner.Scan()

		if err := cmds.Execute(scanner.Text()); err != nil {
			fmt.Println(err.Error())
		}
	}
}