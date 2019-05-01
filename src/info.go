package src

import (
	"io/ioutil"
)

// GetCommands .
func GetCommands() ([]string, error) {
	files, err := ioutil.ReadDir("./scripts")
	if err != nil {
		return nil, err
	}

	result := make([]string, len(files))
	for i, file := range files {
		result[i] = file.Name()
	}

	return result, nil
}
