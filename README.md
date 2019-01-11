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
- command: prints a `:` so that it is clear that the user can enter a command

## Install instructions

Download and install:
```bash
go get github.com/cezarmathe/gosh
```
