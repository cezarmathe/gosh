package prompt

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

// Cmdprompt holds all key information for the prompt.
type Cmdprompt struct {
	username string
	hostname string
	prompt   string
}

// New returns a Cmdprompt pre-populated with various system-level information
// such as the username and hostname
// the string representation of it returns a prompt with the current
// working directory as well as the cached username and host.
func New(promptChar string) *Cmdprompt {
	if len(promptChar) == 0 {
		promptChar = ":"
	}
	return &Cmdprompt{
		username: username(),
		hostname: hostname(),
		prompt:   promptChar,
	}
}

func (i *Cmdprompt) String() string {
	cwd := workdir()
	return fmt.Sprintf("%s@%s %s %s ", i.username, i.hostname, cwd, i.prompt)
}

func username() string {
	fallbackUsername := "unknown" // this should be rare.
	user, err := user.Current()
	if err != nil {
		return fallbackUsername
	}
	return user.Username
}

func hostname() string {
	fallbackHostname := "localhost"
	hn, err := os.Hostname()
	if err != nil {
		return fallbackHostname
	}
	return hn
}

func workdir() string {
	fallbackWorkdir := "?" // this should be rare.
	wd, err := os.Getwd()
	if err != nil {
		wd = fallbackWorkdir
	}

	homedir := os.Getenv("HOME")
	if homedir == wd {
		return "~"
	} else if strings.HasPrefix(wd, homedir) {
		return "~" + strings.TrimPrefix(wd, homedir)
	} else {
		return wd
	}
}
