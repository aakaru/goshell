package main

import (
	"github.com/c-bata/go-prompt"
	"os"
	"path/filepath"
	"strings"
)

func completer(d prompt.Document) []prompt.Suggest {
	text := d.TextBeforeCursor()
	args := strings.Fields(text)

	var suggestions []prompt.Suggest

	if len(args) <= 1 {
		for cmd := range builtinCommands {
			suggestions = append(suggestions, prompt.Suggest{
				Text:        cmd,
				Description: "Built-in command",
			})
		}
		pathEnv := os.Getenv("PATH")
		paths := strings.Split(pathEnv, string(os.PathListSeparator))

		for _, path := range paths {
			files, err := os.ReadDir(path)
			if err != nil {
				continue
			}
			for _, file := range files {
				if !file.IsDir() {
					suggestions = append(suggestions, prompt.Suggest{
						Text:        file.Name(),
						Description: "External Command",
					})
				}
			}
		}
	} else {
		prefix := ""
		if len(args) > 0 {
			prefix = args[len(args)-1]
		}

		dir := "."
		if filepath.IsAbs(prefix) {
			dir = filepath.Dir(prefix)
		} else if strings.Contains(prefix, string(os.PathListSeparator)) {
			dir = filepath.Dir(prefix)
		}

		files, err := os.ReadDir(dir)
		if err == nil {
			for _, file := range files {
				name := file.Name()
				if file.IsDir() {
					name += string(os.PathSeparator)
				}
				if strings.HasPrefix(name, filepath.Base(prefix)) {
					path := filepath.Join(filepath.Dir(prefix), name)
					if filepath.Dir(prefix) == "." {
						path = name
					}

					fileType := "File"
					if file.IsDir() {
						fileType = "Directory"
					}
					suggestions = append(suggestions, prompt.Suggest{
						Text:        path,
						Description: fileType,
					})
				}
			}
		}
	}
	return suggestions
}
