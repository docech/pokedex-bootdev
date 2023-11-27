package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/docech/pokedex-bootdev/api/http"
	"github.com/docech/pokedex-bootdev/api/pokeapi"
	"github.com/docech/pokedex-bootdev/cli/commands"
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
	cmds.Register(commands.NewMapCommand(
		locationAreaResource.Next,
		locationAreaResource.Data,
	))
	cmds.Register(commands.NewMapbCommand(
		locationAreaResource.Previous,
		locationAreaResource.Data,
	))
	cmds.Register(commands.NewExploreCommand())
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