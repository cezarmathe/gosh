package shell

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cezarmathe/gosh/builtin/promptitems"
	"github.com/cezarmathe/gosh/shell/prompt"
)

const (
	shellLoginText = "Hello, %user%"
)

func initialize() {
	fmt.Println(shellLoginText)

	reader = bufio.NewReader(os.Stdin)

	shellPrompt = prompt.Prompt{
		Items: []prompt.Item{
			promptitems.UserItem{},
			promptitems.HostnameItem{},
			promptitems.WorkdirItem{},
			promptitems.CommandpromptItem{},
		},
	}

	shellPrompt.Generate()
}
