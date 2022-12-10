package client

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/urfave/cli/v2"
	client "github.com/wuqinqiang/easycar/client"
	"github.com/wuqinqiang/easycar/core/registry"
	"github.com/wuqinqiang/easycar/core/registry/consulx"
	"github.com/wuqinqiang/easycar/core/registry/etcdx"
	"time"
)

var DiscoveryCmd = &cli.Command{
	Name:    "discovery",
	Aliases: []string{"discovery"},
	Usage:   "connection easycar by discovery",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "mod",
		},
		&cli.IntFlag{
			Name: "count",
		},
	},

	Action: func(cliCtx *cli.Context) error {
		server := cliCtx.String("easycar")

		count := 1
		if cliCtx.Int("count") > 1 {
			count = cliCtx.Int("count")
		}

		var (
			d   registry.Discovery
			err error
		)

		d, err = etcdx.New(etcdx.Conf{
			Hosts: []string{"127.0.0.1:2379"}})
		if err != nil {
			return err
		}
		m := cliCtx.String("mod")
		if m == "consul" {
			client, err := api.NewClient(api.DefaultConfig())
			if err != nil {
				return err
			}
			d = consulx.New(client)
		}

		client.RegisterBuilder(d)

		cli, err := client.New(server, client.WithDiscovery())
		if err != nil {
			return err
		}

		defer func() {
			time.Sleep(3 * time.Minute)
			defer cli.Close(context.Background())
		}()

		for i := 0; i < count; i++ {
			ctx, cancel := context.WithTimeout(cliCtx.Context, 5*time.Second)
			defer cancel()
			gid, err := cli.Begin(ctx)
			if err != nil {
				return err
			}
			fmt.Println("Begin gid:", gid)

			if err = cli.Register(ctx, gid, GetSrv()); err != nil {
				return err
			}

			if err := cli.Start(ctx, gid); err != nil {
				fmt.Println("start err:", err)
			}
			fmt.Println("end gid:", gid)
			time.Sleep(3 * time.Second)
		}

		return nil
	},
}
