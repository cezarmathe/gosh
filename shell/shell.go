package shell

import (
	"bufio"
	"os"

	"github.com/cezarmathe/gosh/process"
	"github.com/cezarmathe/gosh/shell/prompt"
)

var (
	// shellPrompt contains what gets displayed each time the user is prompted for input
	shellPrompt prompt.Prompt
	// reader reads the user input
	reader *bufio.Reader
	// shellFiles contains the files that are open to each process
	shellFiles []*os.File
	// a controller for shell processes
	processController process.ProcessController
)

// Run contains the entire shell lifecycle
func Run() {

	// initialize the shell
	initialize()

	defer terminate()

	// interpret each input indefinetly
	for {
		interpret()
	}
}
