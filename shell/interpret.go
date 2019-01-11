package shell

import (
	"fmt"
	"os"
	"strings"
)

func interpret() {
	// print the prompt
	shellPrompt.Print()

	// read user input
	text, err := readInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	// if no input is given, return
	if text == "" {
		return
	}

	fields := strings.Fields(text)
	if len(fields) == 1 {
		processController.StartProcess(fields[0], []string{}, &os.ProcAttr{Files: shellFiles})
	} else {
		processController.StartProcess(fields[0], fields[1:], &os.ProcAttr{Files: shellFiles})
	}

}

func readInput() (string, error) {
	return reader.ReadString('\n')
}
