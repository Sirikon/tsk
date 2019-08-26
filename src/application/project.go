package application

import (
	"github.com/sirikon/tsk/src/application/filesystem"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

type Project struct {
	TskFile    *TskFile
	Commands   []*Command
	RootFolder string
}

func GetProject(cwd string) (*Project, error) {
	rootFolder, err := filesystem.FindTskProjectRoot(cwd)
	if err != nil {
		return nil, err
	}

	return getProject(rootFolder)
}

func getProject(rootFolder string) (*Project, error) {
	project := &Project{
		RootFolder: rootFolder,
	}

	data, err := ioutil.ReadFile(path.Join(rootFolder, "Tskfile.yml"))
	if err != nil {
		return nil, err
	}

	tskFile, err := parseTskFile(data)
	if err != nil {
		return nil, err
	}

	project.TskFile = tskFile

	commands, err := getCommandsInFolder(path.Join(rootFolder, "scripts"))
	if err != nil {
		return nil, err
	}

	project.Commands = commands

	return project, nil
}

func getCommandsInFolder(folderPath string) ([]*Command, error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	result := make([]*Command, len(files))
	for i, file := range files {
		command := &Command{
			Name:     removeFileExtension(file.Name()),
			Path:     path.Join(folderPath, file.Name()),
			Runnable: !file.IsDir(),
		}

		if file.IsDir() {
			commands, err := getCommandsInFolder(path.Join(folderPath, file.Name()))
			if err != nil {
				return nil, err
			}
			command.SubCommands = commands
		}

		result[i] = command
	}

	return result, nil
}

func removeFileExtension(name string) string {
	return strings.TrimSuffix(name, filepath.Ext(name))
}
