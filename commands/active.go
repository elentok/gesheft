package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
)

var Active = cli.Command{
	Name:      "active",
	ShortName: "a",
	Usage:     "lists the active tunnels",
	Action:    active,
}

func active(c *cli.Context) {

	active, err := config.GetActive()

	if err != nil {
		helpers.ExitWithError(err)
	}

	for name, pid := range active {
		fmt.Printf("%10d %s\n", pid, name)
	}
}
