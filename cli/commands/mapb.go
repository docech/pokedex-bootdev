package commands

type mapbCommand struct {
}

func NewMapbCommand() *mapbCommand {
	return &mapbCommand{}
}

func (c *mapbCommand) Execute() error {
	return nil
}

func (c mapbCommand) About() aboutCommand {
	return aboutCommand{
		name: "mapb",
		description: `Opposite of map command. Displays previous 20 location areas from Pokemon world
			and subsequent calls will display the previous 20 areas.`,
	}
}