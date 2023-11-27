package commands

import (
	"errors"
	"strings"
)

type executableCommand interface {
	Execute(params ...string) error
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

func (c cliCommands) Execute(params string) error { 
	normParams := normParams(strings.Split(params, " "))
	if len(normParams) == 0 {
		return errors.New("missing command name")
	}

	commandName := normParams[0]
	cmd, ok := c.commands[commandName]
	if !ok {
		return errors.New("command not found")
	}

	return cmd.Execute(normParams[1:]...)
}

func normParams(params []string) []string {
	normParams := []string{}
	for _, param := range params {
		normParams = append(normParams, strings.TrimSpace(strings.ToLower(param)))
	}
	return normParams
}
