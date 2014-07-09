package commands

import (
	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/helpers"
	"github.com/elentok/gesheft/tunnel"
)

var Active = cli.Command{
	Name:      "active",
	ShortName: "a",
	Usage:     "lists the active tunnels",
	Action:    active,
}

func active(c *cli.Context) {

	active, err := tunnel.GetActive()

	if err != nil {
		helpers.ExitWithError(err)
	}

	active.Print()
}
