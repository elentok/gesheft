package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
)

var Show = cli.Command{
	Name:      "show",
	ShortName: "s",
	Usage:     "Show tunnel information",
	Action:    show,
}

func show(c *cli.Context) {
	config, err := config.Get()
	if err != nil {
		helpers.ExitWithError(err)
	}

	name := c.Args().First()

	tunnel, ok := config.Tunnels()[name]
	if !ok {
		helpers.ExitWithMessage(fmt.Sprintf("No tunnel named '%s'", name))
	}

	tunnel.Print()
}
