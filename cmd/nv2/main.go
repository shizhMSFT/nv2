package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "nv2",
		Usage: "Notary V2 - Prototype",
		Authors: []*cli.Author{
			{
				Name:  "Shiwei Zhang",
				Email: "shizh@microsoft.com",
			},
			{
				Name:  "Aviral Takkar",
				Email: "avtakkar@microsoft.com",
			},
		},
		Commands: []*cli.Command{
			signCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
