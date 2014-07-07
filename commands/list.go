package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/config"
	"github.com/elentok/gesheft/helpers"
)

var List = cli.Command{
	Name:      "list",
	ShortName: "l",
	Usage:     "list all available tunnels",
	Flags: []cli.Flag{
		cli.BoolFlag{"short, s",
			"print in a single line (useful for shell completions)"},
		cli.BoolFlag{"detailed, d",
			"print the full details of every tunnel"},
	},
	Action: list,
}

func list(c *cli.Context) {
	cfg, err := config.Get()

	if err != nil {
		helpers.ExitWithError(err)
	}

	detailed := c.Bool("detailed")
	short := c.Bool("short")

	for name, tunnel := range cfg.Tunnels() {
		if detailed {
			tunnel.Print()
			fmt.Println()
		} else if short {
			fmt.Printf("%s ", name)
		} else {
			fmt.Println(name)
		}
	}
}
