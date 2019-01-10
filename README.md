# Gosh - **WIP**

A shell written in go

## Features

Gosh includes two builtin commands: 
- `cd <dir_name>` 
  - running `cd` with no arguments changes the directory to $HOME
- `exit <exit_code(optional)>`

Gosh comes with 4 builtin prompt items:
- user: displays the current user
- hostname: displays your hostname
- workdir: displays your current working directory
  - paths relative to $HOME are printed as `~/path`
- command: prints a `:` so that it clears that the user can enter a command

## Install instructions

This installation process assumes you have the Go environment variables properly set up and that you added GOBIN to your PATH.

- `$ git clone https://github.com/cezarmathe/gosh.git`
- `$ cd gosh`
- `$ go install`

After this, you can run `$ gosh` to start the shell.