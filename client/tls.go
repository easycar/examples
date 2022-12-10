package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/urfave/cli/v2"
	client "github.com/wuqinqiang/easycar/client"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"time"
)

var TlsCmd = &cli.Command{
	Name:  "tls",
	Usage: "connection easrcar with tls",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "crtFile",
			Value: "key/ca.crt",
			Usage: "client ca file",
		},
		&cli.StringFlag{
			Name:  "serverName",
			Value: "www.easycar.com",
			Usage: "server name for ca",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		server := cliCtx.String("easycar")
		crtFile := cliCtx.String("crtFile")
		serverName := cliCtx.String("serverName")

		b, err := ioutil.ReadFile(crtFile)
		if err != nil {
			log.Fatal(err)
		}
		cp := x509.NewCertPool()
		if !cp.AppendCertsFromPEM(b) {
			log.Fatal(fmt.Errorf("credentials: failed to append certificates"))
		}
		tls := &tls.Config{
			ServerName: serverName,
			RootCAs:    cp,
		}

		var opts []client.Option
		opts = append(opts, client.WithTls(tls))
		opts = append(opts, client.WithGrpcDailOpts([]grpc.DialOption{grpc.WithBlock(), grpc.WithReturnConnectionError()}))
		opts = append(opts, client.WithConnTimeout(5*time.Second))

		cli, err := client.New(server, opts...)
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
