package main

import (
	"github.com/cezarmathe/gosh/shell"
)

func init() {

}

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Println("GoSH v0.0.1")

	shell.Run()
	// fmt.Print("> ")
	// text, _ := reader.ReadString('\n')

	// command := command.GetCommand(text)

	// command.Execute()

	// if command[0] == "exit" {
	// 	os.Exit(0)
	// } else if command[0] == "cd" {
	// 	os.Chdir(command[1])
	// }

	// commandOutput, _ := exec.Command(command[0], command[1:]...).CombinedOutput()

	// fmt.Printf("%s\n", commandOutput)
}
