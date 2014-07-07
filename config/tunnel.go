package config

import "fmt"

type Tunnel struct {
	Name     string
	Host     string
	Username string
	Port     int

	Bind Bind
}

type Bind struct {
	ClientPort int `yaml:"client-port"`
	HostPort   int `yaml:"host-port"`
	Host       string
}

func (t *Tunnel) Print() {
	fmt.Printf("%s:\n", t.Name)
	fmt.Printf("  %s@%s:%d\n", t.Username, t.Host, t.Port)
	fmt.Printf("  -L %s\n", t.Bind.ToString())
}

func (b *Bind) ToString() string {
	return fmt.Sprintf("%d:%s:%d", b.ClientPort, b.Host, b.HostPort)

}
