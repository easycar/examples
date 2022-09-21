package stock

import (
	"context"
	"flag"
	"fmt"
	"github.com/easycar/examples/srvpb/stock"
	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/grpc"
	pbProto "google.golang.org/protobuf/proto"
	"log"
	"net"
)

var (
	_ stock.StockServer = (*Srv)(nil)
)

type Srv struct {
	stock.UnimplementedStockServer
}

func (s Srv) TryDeduct(ctx context.Context, req *stock.Req) (*stock.TryDeductResp, error) {
	fmt.Printf("[Stock]TryDeduct req:%+v\n", req)
	return new(stock.TryDeductResp), nil
}

func (s Srv) ConfirmDeduct(ctx context.Context, req *stock.Req) (*stock.ConfirmDeductResp, error) {
	fmt.Printf("[Stock]ConfirmDeduct req:%v\n", req)
	return new(stock.ConfirmDeductResp), nil
}

func (s Srv) CancelDeduct(ctx context.Context, req *stock.Req) (*stock.CancelDeductResp, error) {
	fmt.Printf("[Stock]CancelDeuct req:%v\n", req)
	return new(stock.CancelDeductResp), nil
}

func Start(port int) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed listen to:%v", err)
	}
	s := grpc.NewServer()
	stock.RegisterStockServer(s, new(Srv))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to order server:%v", err)
		}
	}()
	fmt.Println("stock server start:", port)
}

func NewData() []byte {
	req := &stock.Req{
		SkuId:  "remember250",
		Number: 10,
	}

	b, _ := pbProto.Marshal(req)
	return b
}

func RegisterTcc(port int) (branches []*proto.RegisterReq_Branch) {
	uri := fmt.Sprintf("127.0.0.1:%d", port)
	reqData := NewData()
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/stock.Stock/TryDeduct",
		ReqData:  string(reqData),
		TranType: proto.TranType_TCC,
		Protocol: "grpc",
		Action:   proto.Action_TRY,
		Level:    1,
	})

	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/stock.Stock/ConfirmDeduct",
		ReqData:  string(reqData),
		TranType: proto.TranType_TCC,
		Protocol: "grpc",
		Action:   proto.Action_CONFIRM,
		Level:    1,
	})

	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/stock.Stock/CancelDeduct",
		ReqData:  string(reqData),
		TranType: proto.TranType_TCC,
		Protocol: "grpc",
		Action:   proto.Action_CANCEL,
		Level:    1,
	})
	return
}
