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
	cmds.Register(commands.NewMapCommand())
	cmds.Register(commands.NewMapbCommand())
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