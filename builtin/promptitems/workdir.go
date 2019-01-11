package promptitems

import (
	"fmt"
	"os"
	"strings"

	"github.com/cezarmathe/gosh/shell/prompt"
)

type WorkdirItem struct {
	Workdir string
}

func (i WorkdirItem) Generate() prompt.Item {
	return i
}

func (i WorkdirItem) Print() {

	wd, _ := os.Getwd()

	if wd == os.Getenv("HOME") {
		fmt.Print("~")
	} else if strings.HasPrefix(wd, os.Getenv("HOME")) {
		fmt.Print("~" + strings.TrimPrefix(wd, os.Getenv("HOME")))
	} else {
		fmt.Print(wd)
	}
}

func (i WorkdirItem) Prefix() {
	fmt.Print(" ")
}
