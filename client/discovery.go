package client

import (
	"context"
	"fmt"
	client "github.com/easycar/client-go"
	"github.com/urfave/cli/v2"
	"github.com/wuqinqiang/easycar/core/registry/etcdx"
	"log"
	"time"
)

var DiscoveryCmd = &cli.Command{
	Name:    "discovery",
	Aliases: []string{"discovery"},
	Usage:   "connection easycar by discovery",
	Action: func(cliCtx *cli.Context) error {
		serverUrl := cliCtx.String("easycar")
		r, err := etcdx.NewRegistry(etcdx.Conf{
			Hosts: []string{"127.0.0.1:2379"}})
		if err != nil {
			log.Fatal(err)
		}

		cli, err := client.New(serverUrl, client.WithDiscovery(r))
		if err != nil {
			log.Fatal(err)
		}

		ctx, cancel := context.WithTimeout(cliCtx.Context, 5*time.Second)
		defer cancel()
		defer cli.Close(ctx)

		gid, err := cli.Begin(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Begin gid:", gid)

		if err = cli.AddGroup(false, GetSrv()...).
			Register(ctx); err != nil {
			log.Fatal(err)
		}

		if err := cli.Start(ctx); err != nil {
			fmt.Println("start err:", err)
		}
		fmt.Println("end gid:", gid)
		return nil
	},
}
