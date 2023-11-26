package commands

type mapCommand struct {
}

func NewMapCommand() *mapCommand {
	return &mapCommand{}
}

func (c *mapCommand) Execute() error {
	return nil
}

func (c mapCommand) About() aboutCommand {
	return aboutCommand{
		name: "map",
		description: `Displays 20 location areas from Pokemon world.
			Subsequent calls will display the next 20 areas.`,
	}
}