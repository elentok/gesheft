package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

type Config interface {
	Tunnels() map[string]Tunnel
}

type configImpl struct {
	tunnels map[string]Tunnel
}

func (c *configImpl) Tunnels() map[string]Tunnel {
	return c.tunnels
}

func Load(filename string) (Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	hash := make(map[string]interface{})

	err = yaml.Unmarshal(bytes, &hash)
	if err != nil {
		return nil, err
	}

	cfg := configImpl{
		tunnels: make(map[string]Tunnel),
	}
	for name, attributes := range hash {
		attribs := attributes.(map[interface{}]interface{})
		tunnel := NewTunnel(name, attribs)
		cfg.tunnels[name] = tunnel
	}

	return &cfg, nil
}

func Get() (Config, error) {
	return Load(getConfigFilepath())
}

func getConfigFilepath() string {
	return filepath.Join(os.Getenv("HOME"), ".shaft")
}
