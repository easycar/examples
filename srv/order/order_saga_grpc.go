package order

import (
	"context"
	"fmt"
	"github.com/easycar/examples/srvpb/order"
	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/grpc"
	pbProto "google.golang.org/protobuf/proto"
	"log"
	"net"
	"time"
)

var (
	_ order.OrderServer = (*Srv)(nil)
)

type Srv struct {
	order.UnimplementedOrderServer
}

func (s Srv) Create(ctx context.Context, req *order.Req) (*order.CreateResp, error) {
	time.Sleep(400 * time.Millisecond)
	fmt.Printf("[Order]create order req userId %v skuId %v \n", req.GetSkuId(), req.GetUserId())
	return new(order.CreateResp), nil
}

func (s Srv) Cancel(ctx context.Context, req *order.Req) (*order.CancelResp, error) {
	fmt.Printf("[Order]Cancel order req :%+v\n", req)
	return new(order.CancelResp), nil
}

func Start(port int) {
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

func NewData() []byte {
	req := &order.Req{
		UserId: "remember",
		SkuId:  "520",
		Amount: 100,
	}

	b, _ := pbProto.Marshal(req)
	return b
}

func RegisterSaga(port int) (branches []*proto.RegisterReq_Branch) {
	uri := fmt.Sprintf("127.0.0.1:%d", port)
	reqData := NewData()

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
