package command

import (
	"os"
	"os/exec"
	"strings"
)

// Command contains the parameters for executing a certain command
type Command struct {
	Exec string
	Args []string
}

// GetCommand returns a command from a string
func GetCommand(text string) Command {

	fields := strings.Fields(text)

	if len(fields) > 1 {
		return Command{
			Exec: fields[0],
			Args: fields[1:],
		}
	}
	return Command{
		Exec: fields[0],
	}
}

// HasArgs returns true if the command has arguments or false if it does not have arguments
func (c Command) HasArgs() bool {
	return len(c.Args) > 0
}

// Execute executes the command
func (c Command) Execute() error {
	var cmd *exec.Cmd

	if c.HasArgs() {
		cmd = exec.Command(c.Exec, c.Args...)
	} else {
		cmd = exec.Command(c.Exec)
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
