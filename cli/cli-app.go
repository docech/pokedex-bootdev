package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/docech/pokedex-bootdev/api"
	"github.com/docech/pokedex-bootdev/cli/internal/commands"
	"github.com/docech/pokedex-bootdev/domain/pokemonworld"
)

type CliApp interface {
	Run() error
}

type cliApp struct {
	commands *commands.CliCommands
}

func InitiateCliApp() CliApp {
	cmds := commands.NewCliCommands()
	
	pokeApi := api.PokeAPI()
	pokedex := pokemonworld.NewPokedex()

	cmds.Register(commands.NewHelpCommand(commands.HelpDeps{
		ProvideAbouts: cmds.About,
	}))
	cmds.Register(commands.NewMapCommand(
		pokeApi.LocationAreasResource.Next,
		pokeApi.LocationAreasResource.Data,
	))
	cmds.Register(commands.NewMapbCommand(
		pokeApi.LocationAreasResource.Previous,
		pokeApi.LocationAreasResource.Data,
	))
	cmds.Register(commands.NewExploreCommand(
		pokeApi.LocationAreaResource.Detail,
		pokeApi.LocationAreaResource.Data,
	))
	cmds.Register(commands.NewCatchCommand(
		pokedex,
		pokeApi.PokemonResource.Detail,
		pokeApi.PokemonResource.Data,
	))
	cmds.Register(commands.NewInspectCommand(
		pokedex,
	))
	cmds.Register(commands.NewPokedexCommand(
		pokedex,
	))
	cmds.Register(commands.NewExitCommand())

	return &cliApp{
		commands: cmds,
	}
}

func (c *cliApp) Run() error {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Starting Pokedex...")
	for {
		fmt.Print("pokedex > ")
		
		scanner.Scan()

		if err := c.commands.Execute(scanner.Text()); err != nil {
			fmt.Println(err.Error())
		}
	}
}

