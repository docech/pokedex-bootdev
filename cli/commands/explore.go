package commands

type exploreCommand struct {
}

func NewExploreCommand() cliCommand {
	return &exploreCommand{}
}

func (c *exploreCommand) Execute() error {
	return nil
}

func (c *exploreCommand) About() aboutCommand {
	return aboutCommand{
		name:        "explore",
		description: "Explore the Pokemon location area",
	}
}