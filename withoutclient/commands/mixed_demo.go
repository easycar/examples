package commands

import (
	"context"
	"fmt"
	"github.com/easycar/examples/conf"
	"github.com/easycar/examples/srv/account"
	"github.com/easycar/examples/srv/order"
	"github.com/easycar/examples/srv/stock"
	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func RunDemo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	easycar := GetEasyCarClient()
	beginResp, err := easycar.Begin(ctx, new(emptypb.Empty))
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

	settings := conf.New()

	branches = append(branches, order.RegisterSaga(settings.OrderPort)...)
	branches = append(branches, account.RegisterTCC(settings.AccountPort)...)
	branches = append(branches, stock.RegisterTcc(settings.StockPort)...)

	registerReq.Branches = append(registerReq.Branches, branches...)

	if _, err = easycar.Register(ctx, &registerReq); err != nil {
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
			if _, err = easycar.Rollback(ctx, &rolbackReq); err != nil {
				err = fmt.Errorf("gid %v Rollback err:%v", beginResp.GetGId(), err)
				return
			}
			return
		}
		var (
			commitReq proto.CommitReq
		)
		commitReq.GId = beginResp.GetGId()
		if _, err = easycar.Commit(ctx, &commitReq); err != nil {
			err = fmt.Errorf("gid %v Commit err:%v", beginResp.GetGId(), err)
		}
		fmt.Println("end gid:", beginResp.GetGId())
	}()
	// phase1
	if _, err = easycar.Start(ctx, &startReq); err != nil {
		err = fmt.Errorf("start err:%v", err)
	}
	return nil
}
