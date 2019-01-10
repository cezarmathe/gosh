package builtin

import (
	"github.com/cezarmathe/gosh/command"
)

func BuiltinCommand(com command.Command) bool {
	switch com.Exec {
	case "cd":
		ExecCd(com)
		return true
	case "exit":
		ExecExit(com)
		return true
	default:
		return false
	}
}
