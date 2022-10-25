package client

import (
	client "github.com/easycar/client-go"
	"github.com/easycar/examples/srv/account"
	"github.com/easycar/examples/srv/order"
	"github.com/easycar/examples/srv/stock"
)

func GetSrv() (groups []*client.Group) {
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
	groups = append(groups, stockSrv, orderSrv, accountSrv)
	return
}
