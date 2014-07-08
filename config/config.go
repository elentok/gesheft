package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/elentok/gesheft/tunnel"
	"gopkg.in/yaml.v1"
)

type Config interface {
	Tunnels() map[string]*tunnel.Tunnel
}

type configImpl struct {
	tunnels map[string]*tunnel.Tunnel
}

func (c *configImpl) Tunnels() map[string]*tunnel.Tunnel {
	return c.tunnels
}

func Load(filename string) (Config, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cfg := configImpl{
		tunnels: make(map[string]*tunnel.Tunnel),
	}

	err = yaml.Unmarshal(bytes, &cfg.tunnels)
	if err != nil {
		return nil, err
	}

	for name, t := range cfg.tunnels {
		t.Name = name
		if t.Binds == nil {
			t.Binds = []*tunnel.Bind{t.Bind}
		}
	}

	return &cfg, nil
}

func Get() (Config, error) {
	return Load(getConfigFilepath())
}

func GetTunnel(name string) (*tunnel.Tunnel, error) {
	cfg, err := Get()
	if err != nil {
		return nil, err
	}

	t, ok := cfg.Tunnels()[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("No tunnel named '%s'", name))
	}

	return t, nil
}

func getConfigFilepath() string {
	return filepath.Join(os.Getenv("HOME"), ".shaft")
}
