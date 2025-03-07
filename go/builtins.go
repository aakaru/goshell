package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func changeDirectory(args []string) error {
	if len(args) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		return os.Chdir(homeDir)
	}
	path := args[1]
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = filepath.Join(homeDir, path[1:])
	}
	return os.Chdir(path)
}

func printWorkingDirectory(args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}
func echoCommand(args []string) error {
	if len(args) > 1 {
		fmt.Println(strings.Join(args[1:], " "))
	} else {
		fmt.Println()
	}
	return nil
}
func typeCommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: type command_name")
	}

	cmdName := args[1]

	if _, ok := builtinCommands[cmdName]; ok {
		fmt.Printf("%s is a built in command\n", cmdName)
		return nil
	}
	cmdPath, err := findCommandInPath(cmdName)
	if err != nil {
		return err
	}
	fmt.Printf("%s is %s\n", cmdName, cmdPath)
	return nil
}
func exitCommand(args []string) error {
	code := 0
	if len(args) > 1 {
		fmt.Scanf(args[1], "%d", &code)
	}
	os.Exit(code)
	return nil
}
