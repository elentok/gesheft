package commands

import (
	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
)

var Stop = cli.Command{
	Name:   "stop",
	Usage:  "Show tunnel information",
	Action: stop,
}

func stop(c *cli.Context) {
	tunnel, err := config.GetTunnel(c.Args().First())
	if err != nil {
		helpers.ExitWithError(err)
	}

	err = tunnel.Stop(true)
	if err != nil {
		helpers.ExitWithError(err)
	}
}
