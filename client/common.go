package client

import (
	client "github.com/easycar/client-go"
	"github.com/easycar/examples/srv/account"
	"github.com/easycar/examples/srv/order"
	"github.com/easycar/examples/srv/stock"
)

func GetSrv() (groups []*client.Group) {
	stockSrv := client.NewTccGroup("grpc://127.0.0.1:50061/stock.Stock/TryDeduct",
		"127.0.0.1:50061/stock.Stock/ConfirmDeduct",
		"127.0.0.1:50061/stock.Stock/CancelDeduct").SetData(stock.NewData()).SetLevel(3)
	orderSrv := client.NewSagaGroup(
		"127.0.0.1:50060/order.Order/Create",
		"127.0.0.1:50060/order.Order/Cancel").SetData(order.NewData()).SetLevel(2)
	accountSrv := client.NewTccGroup(
		"http://127.0.0.1:50062/account/tryDebit",
		"http://127.0.0.1:50062/account/confirmDebit",
		"http://127.0.0.1:50062/account/cancelDebit").SetData(account.NewData()).
		SetTimeout(2).SetLevel(1)
	groups = append(groups, stockSrv, orderSrv, accountSrv)
	return
}
