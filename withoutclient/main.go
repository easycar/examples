package main

import (
	"flag"
	"github.com/urfave/cli/v2"
	"github.com/wuqinqiang/easycar/examples/withoutclient/commands"
	"github.com/wuqinqiang/easycar/examples/withoutclient/srv/account"
	"github.com/wuqinqiang/easycar/examples/withoutclient/srv/order"
	"github.com/wuqinqiang/easycar/examples/withoutclient/srv/stock"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	orderPort   = flag.Int("orderPort", 50060, "the order sever port")
	stockPort   = flag.Int("stockPort", 50061, "the stock sever port")
	accountPort = flag.Int("accountPort", 50062, " the account server port")
	easyCarAddr = flag.String("addr", "localhost:8089", "the address to connect easycar server")
)

func main() {

	flag.Parse()
	order.Start(*orderPort)
	stock.Start(*stockPort)
	account.Start(*accountPort)
	time.Sleep(500 * time.Millisecond)

	commands.MustLoad(*easyCarAddr)

	runCommands()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-c
}

func runCommands() {
	app := cli.App{
		Commands: []*cli.Command{
			commands.BaseDemo,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
