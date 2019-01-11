package builtin

import (
	"github.com/cezarmathe/gosh/command"
)

func ExecCd(com command.Command) {
	// if len(com.Args) == 0 {
	// 	os.Chdir(os.Getenv("HOME"))
	// } else if com.Args[0][:1] == "/" {
	// 	os.Chdir(com.Args[0])
	// } else {
	// 	dir, _ := os.Getwd()
	// 	os.Chdir(dir + "/" + com.Args[0])
	// }
}
