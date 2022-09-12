package order

import (
	"context"
	"fmt"
	"github.com/wuqinqiang/easycar/examples/srvpb/order"
	"github.com/wuqinqiang/easycar/proto"
	pbProto "google.golang.org/protobuf/proto"

	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	_ order.OrderServer = (*Srv)(nil)
	p int
)

type Srv struct {
	order.UnimplementedOrderServer
}

func (s Srv) Create(ctx context.Context, req *order.Req) (*order.CreateResp, error) {
	fmt.Printf("create order req:%+v\n", req)
	return new(order.CreateResp), nil
}

func (s Srv) Cancel(ctx context.Context, req *order.Req) (*order.CancelResp, error) {
	fmt.Printf("CancelCreate order req:%+v\n", req)
	return new(order.CancelResp), nil
}

func Start(port int) {
	p = port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed listen to:%v", err)
	}
	s := grpc.NewServer()
	order.RegisterOrderServer(s, new(Srv))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to order server:%v", err)
		}
	}()
	fmt.Println("order server start:", port)
}

func RegisterSaga() (branches []*proto.RegisterReq_Branch) {
	uri := fmt.Sprintf("127.0.0.1:%d", p)
	req := &order.Req{
		UserId: "remember",
		SkuId:  "520",
		Amount: 100,
	}

	reqData, _ := pbProto.Marshal(req)

	//SAGE NORMAL
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/order.Order/Create",
		ReqData:  string(reqData),
		TranType: proto.TranType_SAGE,
		Protocol: "grpc",
		Action:   proto.Action_NORMAL,
		Level:    1,
	})

	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/order.Order/Cancel",
		ReqData:  string(reqData),
		TranType: proto.TranType_SAGE,
		Protocol: "grpc",
		Action:   proto.Action_COMPENSATION,
		Level:    1,
	})
	return
}
