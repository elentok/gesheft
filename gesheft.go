package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/elentok/gesheft/commands"
)

func main() {
	app := cli.NewApp()

	app.Name = "gesheft"
	app.Version = "0.0.1"
	app.Usage = "SSH Tunnel Manager (based on shaft by n0nick)"

	app.Commands = []cli.Command{
		commands.List,
		commands.Show,
	}

	app.Run(os.Args)
}