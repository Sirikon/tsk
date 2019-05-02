package src

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Command .
type Command struct {
	Name        string
	Subcommands []Command
}

// GetCommands .
func GetCommands() ([]Command, error) {
	return getCommandsInFolder("./scripts")
}

func getCommandsInFolder(path string) ([]Command, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]Command, len(files))
	for i, file := range files {
		command := Command{
			Name: removeFileExtension(file.Name()),
		}
		if file.IsDir() {
			commands, err := getCommandsInFolder(path + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			command.Subcommands = commands
		}
		result[i] = command
	}

	return result, nil
}

func removeFileExtension(name string) string {
	return strings.TrimSuffix(name, filepath.Ext(name))
}
