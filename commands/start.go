package commands

import (
	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
)

var Start = cli.Command{
	Name:      "start",
	ShortName: "s",
	Usage:     "Show tunnel information",
	Action:    start,
}

func start(c *cli.Context) {
	tunnel, err := config.GetTunnel(c.Args().First())
	if err != nil {
		helpers.ExitWithError(err)
	}

	err = tunnel.Start(true)
	if err != nil {
		helpers.ExitWithError(err)
	}
}
