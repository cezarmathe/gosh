package builtin

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cezarmathe/gosh/command"
)

func ExecExit(com command.Command) {
	if len(com.Args) == 0 {
		os.Exit(0)
	} else if val, err := strconv.Atoi(com.Args[0]); err == nil {
		os.Exit(val)
	} else {
		fmt.Println("Exit codes are integers.")
	}
}
