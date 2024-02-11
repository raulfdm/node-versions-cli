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
		Name:  "node-versions",
		Usage: "A simple CLI to check node versions",
		UsageText: `node-versions all
node-versions lts
node-versions lts --all
node-versions latest
node-versions latest 14
		`,
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
						Name:    "all",
						Usage:   "show all LTS versions",
						Aliases: []string{"a"},
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
				UsageText: `node-versions latest
node-versions latest [major-version]`,
				Action: func(ctx *cli.Context) error {
					if ctx.Args().Len() > 0 {
						desiredMajorVersion := ctx.Args().First()
						version, err := nodeVersions.GetLatestOf(desiredMajorVersion)

						if err != nil {
							return err
						}

						fmt.Println(*version)
					} else {
						fmt.Println(nodeVersions.GetLatest())
					}

					return nil
				},
			},
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
