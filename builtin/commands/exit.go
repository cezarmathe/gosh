package commands

import (
	"errors"
	"os"
	"strconv"
)

// Exit quits the shell
func Exit(argv []string) error {
	if len(argv) == 1 {
		os.Exit(0)
	} else if val, err := strconv.Atoi(argv[1]); err == nil {
		os.Exit(val)
	} else {
		return errors.New("exit: exit codes are integers")
	}
	return nil
}
