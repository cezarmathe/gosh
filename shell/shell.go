package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/cezarmathe/gosh/shell/prompt"
)

var (
	// shellPrompt contains what gets displayed each time the user is prompted for input
	shellPrompt prompt.Prompt
	// reader reads the user input
	reader *bufio.Reader
)

// Run contains the entire shell lifecycle
func Run() {
	// initialize the shell
	initialize()

	// interpret each input indefinetly
	for {
		err := interpret()
		if err == nil {
			continue
		}
		if err == io.EOF {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		fmt.Println(err)
	}

}
