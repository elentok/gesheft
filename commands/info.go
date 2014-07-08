package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
)

var Info = cli.Command{
	Name:      "info",
	ShortName: "i",
	Usage:     "Show tunnel information",
	Action:    info,
}

func info(c *cli.Context) {
	tunnel, err := config.GetTunnel(c.Args().First())
	if err != nil {
		helpers.ExitWithError(err)
	}

	tunnel.Print()

	active, err := tunnel.IsActive()
	if err != nil {
		helpers.ExitWithError(err)
	}

	if active {
		fmt.Println("  [ACTIVE]")
	}
}
