package shell

import (
	"strings"

	"github.com/cezarmathe/gosh/builtin"
	"github.com/cezarmathe/gosh/command"
)

func interpret(cmd string) error {
	cmd = strings.TrimRight(cmd, "\r\n")
	// no command was actually issued.
	if cmd == "" {
		return nil
	}

	com := command.GetCommand(cmd)

	if builtin.BuiltinCommand(com) {
		return nil
	}

	return com.Execute()
}
