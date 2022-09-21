package main

import (
	"context"
	"flag"
	"fmt"
	client "github.com/easycar/client-go"
	"github.com/easycar/examples/srv/account"
	"github.com/easycar/examples/srv/order"
	"github.com/easycar/examples/srv/stock"
	"log"
)

var (
	easyCarAddr = flag.String("addr", "localhost:8089", "the address to connect easycar server")
)

func main() {
	cli := client.New(*easyCarAddr)
	ctx := context.Background()
	defer cli.Close(ctx)

	gid, err := cli.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Begin gid:", gid)

	stockSrv := client.NewTccGroup(client.GRPC, "127.0.0.1:50061/stock.Stock/TryDeduct",
		"127.0.0.1:50061/stock.Stock/ConfirmDeduct",
		"127.0.0.1:50061/stock.Stock/CancelDeduct").SetData(stock.NewData())
	orderSrv := client.NewSagaGroup(client.GRPC,
		"127.0.0.1:50060/order.Order/Create",
		"127.0.0.1:50060/order.Order/Cancel").SetData(order.NewData())
	accountSrv := client.NewTccGroup(client.HTTP,
		"http://127.0.0.1:50062/account/tryDebit",
		"http://127.0.0.1:50062/account/confirmDebit",
		"http://127.0.0.1:50062/account/cancelDebit").SetData(account.NewData())
	if err = cli.AddGroup(false, stockSrv, orderSrv, accountSrv).
		Register(ctx); err != nil {
		log.Fatal(err)
	}

	if err := cli.Start(ctx); err != nil {
		fmt.Println("start err:", err)
	}
	fmt.Println("end gid:", gid)
}
