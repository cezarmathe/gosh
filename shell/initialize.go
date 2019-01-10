package shell

import "fmt"

const (
	shellLoginText = "Hello, %user%"
)

func initialize() {
	fmt.Println(shellLoginText)
}
