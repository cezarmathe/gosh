package shell

import (
	"fmt"

	"github.com/cezarmathe/gosh/builtin"
	"github.com/cezarmathe/gosh/command"
)

func interpret() {

	shellPrompt.Print()

	text, err := readInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	com := command.GetCommand(text)

	if builtin.BuiltinCommand(com) {
		return
	}

	com.Execute()

}

func readInput() (string, error) {
	return reader.ReadString('\n')
}
