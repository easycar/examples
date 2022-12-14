package client

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	client "github.com/wuqinqiang/easycar/client"
	"google.golang.org/grpc"
	"log"
	"time"
)

var DirectCmd = &cli.Command{
	Name:    "direct",
	Aliases: []string{"direct"},
	Usage:   "connection easrcar direct",
	Action: func(cliCtx *cli.Context) error {
		serverUrl := cliCtx.String("easycar")
		var opts []client.Option
		opts = append(opts, client.WithGrpcDailOpts([]grpc.DialOption{grpc.WithBlock(), grpc.WithReturnConnectionError()}))
		opts = append(opts, client.WithConnTimeout(5*time.Second))

		cli, err := client.New(serverUrl, opts...)
		if err != nil {
			log.Fatal(err)
		}
		ctx := context.Background()
		defer cli.Close(ctx)

		gid, err := cli.Begin(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Begin gid:", gid)
		if err = cli.Register(ctx, gid, GetSrv()); err != nil {
			log.Fatal(err)
		}
		if err := cli.Start(ctx, gid); err != nil {
			fmt.Println("start err:", err)
		}
		fmt.Println("end gid:", gid)

		return nil

	},
}
