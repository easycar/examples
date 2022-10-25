package withoutclient

import (
	"github.com/easycar/examples/withoutclient/commands"
	"github.com/urfave/cli/v2"
)

var HttpCmd = &cli.Command{
	Name:    "http",
	Aliases: []string{"http"},
	Usage:   "just request easycar service by http",
	Action: func(cliCtx *cli.Context) error {
		serverUrl := cliCtx.String("easycar")
		commands.MustLoad(serverUrl)
		if err := commands.RunDemo(); err != nil {
			return err
		}
		return nil
	},
}
