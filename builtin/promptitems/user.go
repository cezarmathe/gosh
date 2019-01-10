package promptitems

import (
	"fmt"
	"os/user"

	"github.com/cezarmathe/gosh/shell/prompt"
)

type UserItem struct {
	Username string
}

func (i UserItem) Generate() prompt.Item {
	user, _ := user.Current()
	i.Username = user.Username
	return i
}

func (i UserItem) Print() {
	fmt.Print(i.Username)
}

func (i UserItem) PrintPrefix() {
	fmt.Print(" ")
}
