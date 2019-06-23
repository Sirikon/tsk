package info

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// TskFile .
type TskFile struct {
	Name string
	Env  map[string]string
	CWD  string
}

// BuildEnvVars .
func (t *TskFile) BuildEnvVars() []string {
	result := make([]string, 0)
	for key, value := range t.Env {
		result = append(result, key+"="+value)
	}
	return result
}

// ReadTskFile .
func ReadTskFile(cwd string) (*TskFile, error) {
	tskfileDirectory, err := findTskfileDirectory(cwd)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(path.Join(tskfileDirectory, "Tskfile.yml"))
	if err != nil {
		return nil, err
	}

	tskFile := &TskFile{}
	err = yaml.Unmarshal(data, tskFile)
	if err != nil {
		return nil, err
	}

	tskFile.CWD = tskfileDirectory

	return tskFile, nil
}

func findTskfileDirectory(basepath string) (string, error) {
	rootPath, err := filepath.Abs("/")
	if err != nil {
		return "", err
	}

	fullPath := path.Join(basepath, "Tskfile.yml")
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		if rootPath == basepath {
			return "", errors.New("Tskfile not found")
		}
		return findTskfileDirectory(path.Join(basepath, ".."))
	}
	return basepath, nil
}
