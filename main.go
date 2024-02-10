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
		Commands:    []*cli.Command{},
		// Flags: []cli.Flag{
		// 	&cli.BoolFlag{
		// 		Name:  "all",
		// 		Usage: "show all node versions",
		// 	},
		// 	&cli.BoolFlag{
		// 		Name:  "all-lts",
		// 		Usage: "show all LTS versions",
		// 	},
		// 	&cli.BoolFlag{
		// 		Name:  "lts",
		// 		Usage: "show current LTS version",
		// 	},
		// 	&cli.BoolFlag{
		// 		Name:  "latest",
		// 		Usage: "show latest version",
		// 	},
		// 	&cli.StringFlag{
		// 		Name:  "latest-of",
		// 		Usage: "show latest version of a specific version",
		// 	},
		// },
		Action: func(cCtx *cli.Context) error {
			name := "Nefertiti"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}

			if cCtx.Bool("lts") {
				fmt.Println("Hola", name)
			} else {
				cCtx.App.Command("help").Run(cCtx)
			}

			return nil
		},
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
