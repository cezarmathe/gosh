package command

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/kr/pty"
)

// Command contains the parameters for executing a certain command
type Command struct {
	Exec string
	Args []string
}

// GetCommand returns a command from a string
func GetCommand(text string) Command {

	fields := strings.Fields(text)

	if len(fields) == 0 {
		return Command{}
	} else if len(fields) == 1 {
		return Command{
			Exec: fields[0],
		}
	} else {
		return Command{
			Exec: fields[0],
			Args: fields[1:],
		}
	}
}

// HasArgs returns true if the command has arguments or false if it does not have arguments
func (c Command) HasArgs() bool {
	if len(c.Args) == 0 {
		return false
	}
	return true
}

// Execute executes the command
func (c Command) Execute() {

	var cmd *exec.Cmd

	if c.HasArgs() {
		cmd = exec.Command(c.Exec, c.Args...)
	} else if c.Exec == "" {
		return
	} else {
		cmd = exec.Command(c.Exec)
	}

	in, err := pty.Start(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, in)
}
