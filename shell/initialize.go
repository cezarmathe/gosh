package shell

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cezarmathe/gosh/process"

	"github.com/cezarmathe/gosh/builtin/promptitems"
	"github.com/cezarmathe/gosh/shell/prompt"
)

const (
	// shellLoginText contains that gets printed when the shell is initialized
	shellLoginText = "Hello, %user%"
)

func initialize() {
	fmt.Println(shellLoginText)

	// initialize the process controller
	processController = process.NewProcessController()

	// load the default process files
	shellFiles = make([]*os.File, 3)
	shellFiles[0] = os.Stdin
	shellFiles[1] = os.Stdout
	shellFiles[2] = os.Stderr

	// initialize the reader
	reader = bufio.NewReader(os.Stdin)

	// load the prompt with prompt items
	shellPrompt = prompt.Prompt{
		Items: []prompt.Item{
			promptitems.UserItem{},
			promptitems.HostnameItem{},
			promptitems.WorkdirItem{},
			promptitems.CommandpromptItem{},
		},
	}

	// generate the content of each prompt item
	shellPrompt.Generate()
}
