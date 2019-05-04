package info

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

// Command .
type Command struct {
	Name        string
	Path        string
	Subcommands []*Command
}

// IsRunnable .
func (c *Command) IsRunnable() bool {
	return len(c.Subcommands) == 0
}

// GetCommands .
func GetCommands(tskfile *TskFile) ([]*Command, error) {
	return getCommandsInFolder(path.Join(tskfile.CWD, "scripts"))
}

func getCommandsInFolder(path string) ([]*Command, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	result := make([]*Command, len(files))
	for i, file := range files {
		command := &Command{
			Name: removeFileExtension(file.Name()),
			Path: path + "/" + file.Name(),
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
