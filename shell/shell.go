package shell

import (
	"bufio"

	"github.com/cezarmathe/gosh/shell/prompt"
)

var (
	shellPrompt prompt.Prompt
	reader      *bufio.Reader
)

// Run contains the entire shell lifecycle
func Run() {

	initialize()

	defer terminate()

	for {
		interpret()
	}
}
