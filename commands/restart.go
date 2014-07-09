package commands

import (
	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
	"github.com/elentok/gesheft/tunnel"
)

var Restart = cli.Command{
	Name:      "restart",
	ShortName: "r",
	Usage:     "Restarts the tunnel",
	Action:    restart,
}

func restart(c *cli.Context) {
	t, err := config.GetTunnel(c.Args().First())
	if err != nil {
		helpers.ExitWithError(err)
	}

	active, err := tunnel.GetActive()
	if err != nil {
		helpers.ExitWithError(err)
	}

	if active.IsActive(t.Name) {
		err = t.Stop(true)
		if err != nil {
			helpers.ExitWithError(err)
		}
	}

	err = t.Start(true)
	if err != nil {
		helpers.ExitWithError(err)
	}
}
