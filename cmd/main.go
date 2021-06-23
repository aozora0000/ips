package main

import (
	"fmt"
	"github.com/aozora0000/ips"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	version = "local"
)

func main() {
	app := &cli.App{
		Name:     "ips",
		Usage:    "alias subcommand from file",
		Commands: []*cli.Command{
			&cli.Command{
				Name: "path",
				Usage: "display ips config path",
				Action: ips.Path,
			},
			&cli.Command{
				Name: "edit",
				Usage: "edit ips config",
				Action: ips.Edit,
			},

			&cli.Command{
				Name: "init",
				Usage: "initialize ips config",
				Action: ips.Initialize,
			},
			&cli.Command{
				Name: "get",
				Usage: "get ips",
				Action: ips.Get,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "noheader",
						Aliases: []string{"n"},
					},
				},
			},
		},
		Version:  version,
		Action: ips.Get,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}