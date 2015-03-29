package main

import (
	"os"

	"github.com/codegangsta/cli"
)

type karnCli struct {
	Context *karn
}

func (c *karnCli) init() {
	app := cli.NewApp()
	app.Name = "karn"
	app.Usage = "manage multiple Git identities"
	app.Author = "Adnan Abdulhussein"
	app.Email = "adnan@prydoni.us"
	app.Version = "0.0.1"
	app.Commands = c.commands()
	app.Run(os.Args)
}

func (c *karnCli) commands() []cli.Command {
	return []cli.Command{
		{
			Name:  "update",
			Usage: "Update the current repository with a karn configured identity",
			Action: func(ctx *cli.Context) {
				c.Context.Update()
			},
		},
		{
			Name:  "init",
			Usage: "Initialise karn for use in a bash compatible shell",
			Action: func(ctx *cli.Context) {
				c.Context.Init()
			},
		},
		{
			Name:  "install",
			Usage: "Install karn for shell, and a sample configuration",
			Action: func(ctx *cli.Context) {
				c.Context.Install()
			},
		},
	}
}
