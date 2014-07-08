package tunnel

import "fmt"

type Bind struct {
	ClientPort int `yaml:"client-port"`
	HostPort   int `yaml:"host-port"`
	Host       string
}

func (b *Bind) ToString() string {
	return fmt.Sprintf("%d:%s:%d", b.ClientPort, b.Host, b.HostPort)

}
