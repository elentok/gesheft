package commands

import (
	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
)

var Restart = cli.Command{
	Name:      "restart",
	ShortName: "r",
	Usage:     "Restarts the tunnel",
	Action:    restart,
}

func restart(c *cli.Context) {
	tunnel, err := config.GetTunnel(c.Args().First())
	if err != nil {
		helpers.ExitWithError(err)
	}

	active, err := tunnel.IsActive()
	if err != nil {
		helpers.ExitWithError(err)
	}

	if active {
		err = tunnel.Stop(true)
		if err != nil {
			helpers.ExitWithError(err)
		}
	}

	err = tunnel.Start(true)
	if err != nil {
		helpers.ExitWithError(err)
	}
}
