package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func getPathEnv() []string {
	pathEnv := os.Getenv("PATH")
	seperator := ":"
	if runtime.GOOS == "windows" {
		seperator = ";"
	}
	return strings.Split(pathEnv, seperator)
}

func getExecutableExtensions() []string {
	if runtime.GOOS == "windows" {
		return []string{"", ".exe", ".cmd", ".bat"}
	}
	return []string{""}
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	if runtime.GOOS == "windows" {
		ext := strings.ToLower(filepath.Ext(path))
		return ext == ".exe" || ext == ".cmd" || ext == ".bat"
	}
	return info.Mode()&0111 != 0
}

func getUserHomeDir() (string, error) {
	return os.UserHomeDir()
}
