package config

type Tunnel interface {
	Name() string
}

type tunnel struct {
	name string
}

func (t *tunnel) Name() string {
	return t.name
}

func NewTunnel(name string, attributes map[interface{}]interface{}) Tunnel {
	tunnel := tunnel{name: name}

	return &tunnel
}
