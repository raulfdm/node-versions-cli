package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := (&cli.App{
		Name:        "node-versions",
		Description: "A simple CLI to check node versions",
		Commands: []*cli.Command{{
			Name:  "all",
			Usage: "show all versions",
			Action: func(cCtx *cli.Context) error {
				fmt.Println("added task: ", cCtx.Args().First())
				return nil
			},
		},
			{
				Name:  "lts",
				Usage: "show LTS version",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "all",
						Usage: "show all LTS versions",
					},
				},
			},
			{
				Name:  "latest",
				Usage: "show latest version",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
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
		Action: func(cCtx *cli.Context) error {

			cCtx.App.Command("help").Run(cCtx)

			return nil
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
