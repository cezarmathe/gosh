package command

import (
	"fmt"
	"os"
	"strings"
)

// Command contains the parameters for executing a certain command
type Command struct {
	Name string
	Argv []string
}

// GetCommand returns a command from a string
func GetCommand(text string) Command {

	fields := strings.Fields(text)

	if len(fields) == 0 {
		return Command{}
	} else if len(fields) == 1 {
		return Command{
			Name: fields[0],
			Argv: []string{fields[0]},
		}
	} else {
		return Command{
			Name: fields[0],
			Argv: fields,
		}
	}
}

// HasArgs returns true if the command has arguments or false if it does not have arguments
// func (c Command) HasArgs() bool {
// 	if len(c.Argv) == 0 {
// 		return false
// 	}
// 	return true
// }

// Execute executes the command
func (c Command) Execute(files []*os.File) {

	process, err := os.StartProcess(c.Name, c.Argv, &os.ProcAttr{
		Files: files,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = process.Wait()

	if err != nil {
		fmt.Println(err)
	}

	// var cmd *exec.Cmd

	// if c.HasArgs() {
	// 	cmd = exec.Command(c.Exec, c.Args...)
	// } else if c.Exec == "" {
	// 	return
	// } else {
	// 	cmd = exec.Command(c.Exec)
	// }

	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// cmd.Run()

}
