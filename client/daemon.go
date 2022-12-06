package client

import (
	"context"
	"fmt"
	client "github.com/easycar/client-go"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"math/rand"
	"sync"
	"time"
)

var DaemonCmd = &cli.Command{
	Name:    "daemon",
	Aliases: []string{"daemon"},
	Action: func(cliCtx *cli.Context) error {
		serverUrl := cliCtx.String("easycar")
		var opts []client.Option
		opts = append(opts, client.WithGrpcDailOpts([]grpc.DialOption{grpc.WithBlock(), grpc.WithReturnConnectionError()}))
		opts = append(opts, client.WithConnTimeout(5*time.Second))

		for {
			random := rand.Intn(8)
			if random == 0 {
				random = 1
			}
			fmt.Println("[client] random:", random)

			var wg sync.WaitGroup
			for i := 0; i < random; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					ctx := context.Background()

					cli, err := client.New(serverUrl, opts...)
					if err != nil {
						fmt.Printf("[client] wrong new err:%v\n", err)
						return
					}
					defer cli.Close(ctx)

					gid, err := cli.Begin(ctx)
					if err != nil {
						fmt.Printf("[client] wrong Begin:%v\n", err)
						return
					}
					fmt.Println("Begin gid:", gid)
					if err = cli.AddGroup(false, GetSrv()...).
						Register(ctx); err != nil {
						fmt.Printf("[client] wrong AddGroup gid:%s err:%v\n", gid, err)
						return
					}

					if err := cli.Start(ctx); err != nil {
						fmt.Printf("[client] wrong Start gid:%s err:%v\n", gid, err)
						return
					}
					fmt.Println("end gid:", gid)
				}()

			}
			wg.Wait()
			time.Sleep(5 * time.Second)
		}

	},
}
