package builtin

import (
	"errors"

	"github.com/cezarmathe/gosh/builtin/commands"
)

// Check verifies if a command is a builtin command
func Check(argv []string) (func([]string) error, error) {
	switch argv[0] {
	case "cd":
		return commands.Cd, nil
	case "exit":
		return commands.Exit, nil
	default:
		return nil, errors.New(argv[0] + " is not a builtin command")
	}
}
