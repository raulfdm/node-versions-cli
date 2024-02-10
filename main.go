package main

import (
	"fmt"
	"log"
	"node-versions-cli/api"
	"node-versions-cli/data"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var nodeVersions *data.NodeVersions

	app := (&cli.App{
		Name:                 "node-versions",
		Version:              "2.0.0",
		Description:          "A simple CLI to check node versions",
		EnableBashCompletion: true,
		Before: func(ctx *cli.Context) error {
			// We don't want to call the API if there's no subcommand
			if ctx.Args().Len() > 0 {
				versions, err := api.GetNodeVersions()

				if err != nil {
					return err
				}

				nodeVersions = versions
			}

			return nil
		},
		Commands: []*cli.Command{{
			Name:  "all",
			Usage: "show all versions",
			Action: func(ctx *cli.Context) error {

				for _, version := range nodeVersions.GetAll() {
					fmt.Println(version)
				}

				return nil
			},
		},
			{
				Name:  "lts",
				Usage: "show LTS version",
				Action: func(ctx *cli.Context) error {

					// If we don't validate the flag, both Actions will be executed
					// and overlap the output
					if !ctx.Bool("all") {
						fmt.Println(nodeVersions.GetCurrentLts())
					}

					return nil
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "all",
						Usage: "show all LTS versions",
						Action: func(ctx *cli.Context, value bool) error {

							for _, version := range nodeVersions.GetAllLts() {
								fmt.Println(version)
							}

							return nil
						},
					},
				},
			},
			{
				Name:  "latest",
				Usage: "show latest version",
				Action: func(ctx *cli.Context) error {
					fmt.Println(nodeVersions.GetLatest())

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "of",
						Usage: "show latest version of a specific version",
					},
				},
			},
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
