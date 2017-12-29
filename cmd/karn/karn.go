package main

import (
	"log"
	"os"

	"github.com/prydonius/karn"
	"github.com/urfave/cli"
)

func main() {
	log.SetFlags(0)
	app := cli.NewApp()
	app.Name = "karn"
	app.Usage = "manage multiple Git identities"
	app.Author = "Adnan Abdulhussein"
	app.Email = "adnan@prydoni.us"
	app.Version = "0.0.4"
	app.Commands = commands()
	app.Run(os.Args)
}

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:  "update",
			Usage: "Update the current repository with a karn configured identity",
			Action: func(ctx *cli.Context) {
				karn.Update()
			},
		},
		{
			Name:  "init",
			Usage: "Initialise karn for use in a bash compatible shell",
			Action: func(ctx *cli.Context) {
				karn.Init()
			},
		},
		{
			Name:  "install",
			Usage: "Install karn for shell, and a sample configuration",
			Action: func(ctx *cli.Context) {
				karn.Install()
			},
		},
	}
}
