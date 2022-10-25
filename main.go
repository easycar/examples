package main

import (
	"github.com/easycar/examples/client"
	"github.com/easycar/examples/withoutclient"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:                 "examples",
		Usage:                "examples for easycar",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			client.DirectCmd,
			client.TlsCmd,
			client.DiscoveryCmd,
			withoutclient.HttpCmd,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "easycar",
				Usage: "set easycar server url",
				Value: "127.0.0.1:8089",
			},
		},
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
