package promptitems

import (
	"fmt"
	"os"

	"github.com/cezarmathe/gosh/shell/prompt"
)

type HostnameItem struct {
	Hostname string
}

func (i HostnameItem) Generate() prompt.Item {
	hostname, _ := os.Hostname()
	i.Hostname = hostname
	return i
}

func (i HostnameItem) Print() {
	fmt.Print("@" + i.Hostname)
}

func (i HostnameItem) PrintPrefix() {
	// fmt.Print(" ")
}
