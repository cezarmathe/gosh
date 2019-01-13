package shell

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/cezarmathe/gosh/builtin"
)

func interpret() error {
	// todo: add signal listener

	// print the prompt
	shellPrompt.Print()

	// read user input
	input, err := readInput()

	// return errors
	if err != nil {
		if err == io.EOF {
			return err
		}
		return err
	}

	input = strings.TrimRight(input, "\r\n")

	// if no input is given, skip the cycle
	if input == "" {
		return nil
	}

	// todo: add history

	// separate the input in arguments
	argv := strings.Fields(input)

	// check if the command is a builtin command
	fn, err := builtin.Check(argv)
	if err == nil {
		err = fn(argv)
		return err
	}

	// otherwise, execute the command
	cmd := exec.Command("bash", "-c", input)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil

}

func readInput() (string, error) {
	return reader.ReadString('\n')
}
