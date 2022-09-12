package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/wuqinqiang/easycar/examples/withoutclient/srv/account"
	"github.com/wuqinqiang/easycar/examples/withoutclient/srv/order"
	"github.com/wuqinqiang/easycar/examples/withoutclient/srv/stock"
	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

var BaseDemo = &cli.Command{
	Name:    "base demo",
	Aliases: []string{"base"},
	Action: func(cli *cli.Context) (err error) {
		easycar := GetEasyCarClient()
		beginResp, err := easycar.Begin(cli.Context, new(emptypb.Empty))
		if err != nil {
			return fmt.Errorf("begin err:%v", err)
		}

		fmt.Println("gid:", beginResp.GetGId())

		var (
			registerReq proto.RegisterReq
		)
		registerReq.GId = beginResp.GetGId()

		var (
			branches []*proto.RegisterReq_Branch
		)

		branches = append(branches, order.RegisterSaga()...)
		branches = append(branches, account.RegisterTCC()...)
		branches = append(branches, stock.RegisterTcc()...)

		registerReq.Branches = append(registerReq.Branches, branches...)

		if _, err = easycar.Register(cli.Context, &registerReq); err != nil {
			return fmt.Errorf("register err:%v", err)
		}
		startReq := proto.StartReq{GId: beginResp.GetGId()}

		defer func() {

			// phase2
			if err != nil {
				var (
					rolbackReq proto.RollBckReq
				)
				rolbackReq.GId = beginResp.GetGId()
				if _, err := easycar.Rollback(cli.Context, &rolbackReq); err != nil {
					err = fmt.Errorf("gid %v Rollback err:%v", beginResp.GetGId(), err)
					return
				}
				return
			}
			var (
				commitReq proto.CommitReq
			)
			commitReq.GId = beginResp.GetGId()
			if _, err := easycar.Commit(cli.Context, &commitReq); err != nil {
				err = fmt.Errorf("gid %v Commit err:%v", beginResp.GetGId(), err)
			}
			return
		}()
		// phase1
		if _, err = easycar.Start(cli.Context, &startReq); err != nil {
			err = fmt.Errorf("start err:%v", err)
		}
		return nil
	},
}
