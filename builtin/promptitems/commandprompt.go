package promptitems

import (
	"fmt"

	"github.com/cezarmathe/gosh/shell/prompt"
)

type CommandpromptItem struct {
	Prompt string
}

func (i CommandpromptItem) Generate() prompt.Item {
	i.Prompt = ": "
	return i
}

func (i CommandpromptItem) Print() {
	fmt.Print(i.Prompt)
}

func (i CommandpromptItem) Prefix() {
	fmt.Print(" ")
}
