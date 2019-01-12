package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/cezarmathe/gosh/prompt"
)

// Run contains the entire shell lifecycle
func Run() {
	reader := bufio.NewReader(os.Stdin)
	shellPrompt := prompt.New(":")
	for {
		fmt.Print(shellPrompt)
		input, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, err)
		}

		if err = interpret(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
