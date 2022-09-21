package main

import (
	"github.com/easycar/examples/conf"
	"github.com/easycar/examples/srv/account"
	"github.com/easycar/examples/srv/order"
	"github.com/easycar/examples/srv/stock"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	settings := conf.New()
	order.Start(settings.OrderPort)
	stock.Start(settings.StockPort)
	account.Start(settings.AccountPort)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-c
}
