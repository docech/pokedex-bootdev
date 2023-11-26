package commands

import "fmt"

type MapResources struct {
	ProvideLocationAreas func () ([]string, error)
}

func printLocationAreas(locationAreas []string) {
	for _, area := range locationAreas {
		fmt.Println(area)
	}
}

func handleExecute(res MapResources) error {
	areas, err := res.ProvideLocationAreas()
	if err != nil {
		return err
	}
	printLocationAreas(areas)
	return nil
}

type mapCommand struct {
	res MapResources
}

func NewMapCommand(res MapResources) *mapCommand {
	return &mapCommand{
		res: res,
	}
}

func (c *mapCommand) Execute() error {
	return handleExecute(c.res)
}

func (c mapCommand) About() aboutCommand {
	return aboutCommand{
		name: "map",
		description: "Displays X location areas from Pokemon world. Subsequent calls will display the next X areas.",
	}
}

type mapbCommand struct {
	res MapResources
}

func NewMapbCommand(res MapResources) *mapbCommand {
	return &mapbCommand{
		res: res,
	}
}

func (c *mapbCommand) Execute() error {
	return handleExecute(c.res)
}

func (c mapbCommand) About() aboutCommand {
	return aboutCommand{
		name: "mapb",
		description: `Opposite of map command. Displays previous X location areas from Pokemon world.`,
	}
}