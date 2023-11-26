package commands

import (
	"errors"
	"strings"
)

type executableCommand interface {
	Execute() error
}

type aboutCommand struct {
	name        string
	description string
}

type descriptiveCommand interface {
	About() aboutCommand
}

type cliCommand interface {
	executableCommand
	descriptiveCommand
}

type cliCommands struct {
	commands map[string]cliCommand
}

func NewCliCommands() *cliCommands {
	return &cliCommands{
		commands: map[string]cliCommand{},
	}
}

func (c *cliCommands) Register(cmd cliCommand) {
	c.commands[cmd.About().name] = cmd
}

func (c cliCommands) About() []aboutCommand {
	abouts := []aboutCommand{}
	for _, cmd := range c.commands {
		abouts = append(abouts, cmd.About())
	}
	return abouts
}

func (c cliCommands) Execute(commandName string) error {
	normCommandName := strings.TrimSpace(strings.ToLower(commandName))
	cmd, ok := c.commands[normCommandName]
	if !ok {
		return errors.New("command not found")
	}

	return cmd.Execute()
}
