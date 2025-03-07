package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func executor(in string) {
	in = strings.TrimSpace(in)
	if in == "" {
		return
	}

	if in == "exit" {
		fmt.Println("Exiting GoShell.")
		os.Exit(0)
	}
	args := strings.Fields(in)
	command := args[0]

	if builtinFunc, ok := builtinCommands[command]; ok {
		err := builtinFunc(args)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		return
	}
	err := executeCommand(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

}

func executeCommand(args []string) error {
	cmdPath, err := findCommandInPath(args[0])
	if err != nil {
		return fmt.Errorf("Command not found: %s", args[0])
	}
	cmd := exec.Command(cmdPath, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func findCommandInPath(command string) (string, error) {
	if strings.Contains(command, string(os.PathSeparator)) {
		if _, err := os.Stat(command); err != nil {
			return command, nil
		}
		return "", fmt.Errorf("Command not found: %s", command)
	}
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))

	var extns []string

	if os.PathSeparator == '\\' {
		extns = []string{"", ".exe", ",cmd", ".bat"}
	} else {
		extns = []string{""}
	}
	for _, path := range paths {
		for _, extn := range extns {
			fullPath := filepath.Join(path, command+extn)
			if _, err := os.Stat(fullPath); err == nil {
				return fullPath, nil
			}
		}
	}
	return "", fmt.Errorf("Command not found: %s", command)
}

func shellPrompt() string {
	pwd, _ := os.Getwd()
	username := os.Getenv("USER")

	if username == "" {
		os.Getenv("USERNAME")
	}
	hostname, _ := os.Hostname()

	return fmt.Sprintf("%s@%s:%s$ ", username, hostname, pwd)
}
