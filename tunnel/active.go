package tunnel

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

type PidByName map[string]int

func LoadActive(filename string) (PidByName, error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	pidByName := make(PidByName)
	err = yaml.Unmarshal(bytes, &pidByName)

	if err != nil {
		return nil, err
	}

	return pidByName, nil
}

func GetActive() (PidByName, error) {
	return LoadActive(getActiveFilepath())
}

func getActiveFilepath() string {
	return filepath.Join(os.Getenv("HOME"), ".shaft.active")
}
