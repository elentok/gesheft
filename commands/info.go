package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
	"github.com/elentok/gesheft/tunnel"
)

var Info = cli.Command{
	Name:      "info",
	ShortName: "i",
	Usage:     "Show tunnel information",
	Action:    info,
}

func info(c *cli.Context) {
	t, err := config.GetTunnel(c.Args().First())
	if err != nil {
		helpers.ExitWithError(err)
	}

	t.Print()

	active, err := tunnel.GetActive()
	if err != nil {
		helpers.ExitWithError(err)
	}

	if active.IsActive(t.Name) {
		fmt.Println("  [ACTIVE]")
	}
}
