package commands

import (
	"os"
	"strings"
)

// Cd changes the shell's directory
func Cd(argv []string) error {
	if len(argv) == 1 {
		os.Chdir(os.Getenv("HOME"))
	} else if argv[1][0:1] == "/" {
		os.Chdir(argv[1])
	} else if argv[1][0:1] == "~" {
		os.Chdir(os.Getenv("HOME") + strings.Join(argv[1:], ""))
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		os.Chdir(wd + "/" + argv[1])
	}
	return nil
}
