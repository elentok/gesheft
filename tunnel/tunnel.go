package tunnel

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type Tunnel struct {
	Name     string
	Host     string
	Username string
	Port     int
	Bind     *Bind
	Binds    []*Bind
}

func (t *Tunnel) Print() {
	fmt.Printf("%s:\n", t.Name)
	fmt.Printf("  %s@%s:%d\n", t.Username, t.Host, t.Port)

	for _, bind := range t.Binds {
		fmt.Printf("  -L %s\n", bind.ToString())
	}
}

func (t *Tunnel) IsActive() (bool, error) {
	pid, err := GetPID(t.Name)
	if err != nil {
		return false, err
	}

	return isProcessLive(pid), nil
}

func isProcessLive(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func (t *Tunnel) Start(verbose bool) error {
	active, err := t.IsActive()
	if err != nil {
		return err
	}

	if active {
		fmt.Printf("Tunnel '%s' is already active.\n", t.Name)
		return nil
	}

	cmd := exec.Command("ssh")
	cmd.Args = t.getSshArgs()

	err = cmd.Start()
	if err != nil {
		return err
	}

	if verbose {
		fmt.Printf("Tunnel started with pid %d\n", cmd.Process.Pid)
	}

	SaveActive(t.Name, cmd.Process.Pid)

	return nil
}

func (t *Tunnel) getSshArgs() []string {
	//"user@host -N -p {port} -L {bind}"
	length := 5 + 2*len(t.Binds)
	args := make([]string, length)

	args[0] = "ssh"
	args[1] = fmt.Sprintf("%s@%s", t.Username, t.Host)
	args[2] = "-N"
	args[3] = "-p"
	args[4] = fmt.Sprintf("%d", t.Port)

	for i, bind := range t.Binds {
		args[5+i] = "-L"
		args[6+i] = bind.ToString()
	}

	return args
}
