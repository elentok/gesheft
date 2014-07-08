package tunnel

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

type PidByName map[string]int

func LoadActive(filename string) (PidByName, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return make(PidByName), nil
	}

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

func SaveActive(tunnelName string, pid int) error {
	filename := getActiveFilepath()
	active, err := LoadActive(filename)
	if err != nil {
		return err
	}

	active[tunnelName] = pid
	bytes, err := yaml.Marshal(active)
	return ioutil.WriteFile(filename, bytes, 0644)
}

func getActiveFilepath() string {
	return filepath.Join(os.Getenv("HOME"), ".shaft.active")
}

func GetPID(name string) (int, error) {
	pidByName, err := GetActive()
	if err != nil {
		return 0, err
	}

	pid, _ := pidByName[name]
	return pid, nil
}
