package client

import (
	"github.com/easycar/examples/srv/account"
	"github.com/easycar/examples/srv/order"
	"github.com/easycar/examples/srv/stock"
	client "github.com/wuqinqiang/easycar/client"
)

func GetSrv() []*client.Group {
	m := client.NewManger()
	m.AddGroup(false, client.NewTccGroup(
		"http://127.0.0.1:50062/account/tryDebit",
		"http://127.0.0.1:50062/account/confirmDebit",
		"http://127.0.0.1:50062/account/cancelDebit").SetData(account.NewData()).
		SetTimeout(2))

	m.AddGroup(true, client.NewSagaGroup(
		"127.0.0.1:50060/order.Order/Create",
		"127.0.0.1:50060/order.Order/Cancel").SetData(order.NewData()))
	m.AddGroup(true, client.NewTccGroup("grpc://127.0.0.1:50061/stock.Stock/TryDeduct",
		"127.0.0.1:50061/stock.Stock/ConfirmDeduct",
		"127.0.0.1:50061/stock.Stock/CancelDeduct").SetData(stock.NewData()))
	return m.Groups()
}
