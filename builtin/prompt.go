package builtin

import (
	"os"
	"os/user"
	"strings"

	"github.com/cezarmathe/gosh/shell/prompt"
)

type BuiltinPromptItem struct {
	Username      string
	Hostname      string
	Commandprompt string
}

func (i BuiltinPromptItem) Generate() prompt.Item {
	user, _ := user.Current()
	i.Username = user.Username
	i.Hostname, _ = os.Hostname()
	i.Commandprompt = ": "
	return i
}

func (i BuiltinPromptItem) String() string {
	wd, _ := os.Getwd()
	if wd == os.Getenv("HOME") {
		wd = "~"
	} else if strings.HasPrefix(wd, os.Getenv("HOME")) {
		wd = "~" + strings.TrimPrefix(wd, os.Getenv("HOME"))
	}
	return strings.Join([]string{
		i.Username,
		"@",
		i.Hostname,
		" ",
		wd,
		i.Commandprompt,
	}, "")
}

func (i BuiltinPromptItem) Prefix() {}

func Prompt() prompt.Prompt {
	return prompt.Prompt{
		Items: []prompt.Item{
			BuiltinPromptItem{},
		},
	}
}
