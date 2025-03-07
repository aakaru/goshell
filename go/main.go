package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

var builtinCommands map[string]func([]string) error

func init() {
	builtinCommands = map[string]func([]string) error{
		"cd":   changeDirectory,
		"pwd":  printWorkingDirectory,
		"echo": echoCommand,
		"type": typeCommand,
		"exit": exitCommand,
	}
}

func main() {
	fmt.Println("GoShell v0.1 - Type 'exit' to exit.")

	p := prompt.New(executor,
		completer,
		prompt.OptionTitle("GoShell"),
		prompt.OptionPrefix(shellPrompt()),
		prompt.OptionLivePrefix(func() (string, bool) {
			return shellPrompt(), true
		}),
	)
	p.Run()

}
