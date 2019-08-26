package application

import (
	"gopkg.in/yaml.v2"
)

// TskFile .
type TskFile struct {
	Name string
	Env  map[string]string
}

func parseTskFile(data []byte) (*TskFile, error) {
	tskFile := &TskFile{}
	err := yaml.Unmarshal(data, tskFile)
	if err != nil {
		return nil, err
	}
	return tskFile, nil
}
