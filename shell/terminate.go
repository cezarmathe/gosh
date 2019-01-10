package shell

import "fmt"

const (
	shellLogoutText = "Bye, %user%"
)

func terminate() {
	fmt.Println(shellLogoutText)
}
