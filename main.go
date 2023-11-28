package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/docech/pokedex-bootdev/api/http"
	"github.com/docech/pokedex-bootdev/api/pokeapi"
	"github.com/docech/pokedex-bootdev/cli/commands"
	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := commands.NewCliCommands()
	caching := http.CacheConfig {
		MaxAge: 60 * time.Second,
	}
	locationAreasResource := pokeapi.NewLocationAreasResource(
		"https://pokeapi.co/api/v2/location-area/", 
		caching,
	)
	locationAreaResource := pokeapi.NewLocationAreaResource(
		"https://pokeapi.co/api/v2/location-area/",
		caching,
	)
	pokemonResource := pokeapi.NewPokemonResource(
		"https://pokeapi.co/api/v2/pokemon/",
		caching,
	)
	
	pokedex := pokemonworld.NewPokedex()

	cmds.Register(commands.NewHelpCommand(commands.HelpDeps{
		ProvideAbouts: cmds.About,
	}))
	cmds.Register(commands.NewMapCommand(
		locationAreasResource.Next,
		locationAreasResource.Data,
	))
	cmds.Register(commands.NewMapbCommand(
		locationAreasResource.Previous,
		locationAreasResource.Data,
	))
	cmds.Register(commands.NewExploreCommand(
		locationAreaResource.Detail,
		locationAreaResource.Data,
	))
	cmds.Register(commands.NewCatchCommand(
		pokedex,
		pokemonResource.Detail,
		pokemonResource.Data,
	))
	cmds.Register(commands.NewInspectCommand(
		pokedex,
	))
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