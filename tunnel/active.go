package tunnel

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"gopkg.in/yaml.v1"
)

type Active interface {
	SetPID(name string, pid int)
	GetPID(name string) (int, bool)
	IsActive(name string) bool
	Kill(name string, verbose bool) error
	Save() error
	Print()
	RemoveZombies() error
}

type active struct {
	filename  string
	pidByName map[string]int
}

func LoadActive(filename string) (Active, error) {
	a := active{
		filename:  filename,
		pidByName: make(map[string]int),
	}

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return &a, nil
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &a.pidByName)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func GetActive() (Active, error) {
	return LoadActive(getActiveFilepath())
}

func (a *active) SetPID(tunnelName string, pid int) {
	a.pidByName[tunnelName] = pid
}

func (a *active) GetPID(name string) (int, bool) {
	pid, ok := a.pidByName[name]
	return pid, ok
}

func (a *active) IsActive(name string) bool {
	pid, ok := a.GetPID(name)
	if ok {
		return isProcessLive(pid)
	}

	return false
}

func isProcessLive(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func (a *active) Kill(name string, verbose bool) error {
	if a.IsActive(name) {
		pid, _ := a.GetPID(name)

		if verbose {
			fmt.Printf("Stopping tunnel '%s' (pid %d)\n", name, pid)
		}

		err := killProcess(pid)
		if err != nil {
			return err
		}

		delete(a.pidByName, name)
		a.Save()
	}

	return nil
}

func killProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = process.Kill()
	time.Sleep(1 * time.Second)
	return err
}

func (a *active) Save() error {
	bytes, err := yaml.Marshal(a.pidByName)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(a.filename, bytes, 0644)
}

func (a *active) Print() {
	for name, pid := range a.pidByName {
		fmt.Printf("%10d %s\n", pid, name)
	}
}

func getActiveFilepath() string {
	return filepath.Join(os.Getenv("HOME"), ".shaft.active")
}

func (a *active) RemoveZombies() error {
	changed := false

	for name, _ := range a.pidByName {
		if !a.IsActive(name) {
			delete(a.pidByName, name)
			fmt.Printf("Removing zombie tunnel '%s'\n", name)
			changed = true
		}
	}

	if changed {
		return a.Save()
	}

	return nil
}
