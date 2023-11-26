package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/docech/pokedex-bootdev/cli/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := commands.NewCliCommands()
	cmds.Register(commands.NewHelpCommand(commands.HelpResources{
		ProvideAbouts: cmds.About,
	}))
	cmds.Register(commands.NewMapCommand(commands.MapResources{
		ProvideLocationAreas: func () ([]string, error) {
			return []string{"Pallet Town", "Route 1", "Viridian City"}, nil
		},
	}))
	cmds.Register(commands.NewMapbCommand(commands.MapResources{
		ProvideLocationAreas: func () ([]string, error) {
			return []string{"Route 2", "Viridian Forest"}, nil
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