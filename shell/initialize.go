package shell

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cezarmathe/gosh/builtin"
)

const (
	// shellLoginText contains that gets printed when the shell is initialized
	shellLoginText = "Gosh v0.1.0\n"
)

func initialize() {
	fmt.Print(shellLoginText)

	// todo: add configuration

	// initialize the reader
	reader = bufio.NewReader(os.Stdin)

	// load the prompt with prompt items
	// todo: load prompt items if the prompt is configured in the config file
	shellPrompt = builtin.Prompt()

	// generate the content of each prompt item
	shellPrompt.Generate()
}
