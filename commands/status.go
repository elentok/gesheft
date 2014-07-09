package commands

import (
	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/helpers"
	"github.com/elentok/gesheft/tunnel"
)

var Status = cli.Command{
	Name:   "status",
	Usage:  "lists the active tunnels",
	Action: active,
}

func active(c *cli.Context) {

	active, err := tunnel.GetActive()
	if err != nil {
		helpers.ExitWithError(err)
	}

	err = active.RemoveZombies()
	if err != nil {
		helpers.ExitWithError(err)
	}

	active.Print()
}
