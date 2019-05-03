package info

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// TskFile .
type TskFile struct {
	Name string
	Env  map[string]string
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
func ReadTskFile() (*TskFile, error) {
	data, err := ioutil.ReadFile("./Tskfile.yml")
	if err != nil {
		return nil, err
	}

	tskFile := &TskFile{}
	err = yaml.Unmarshal(data, tskFile)
	if err != nil {
		return nil, err
	}

	return tskFile, nil
}
