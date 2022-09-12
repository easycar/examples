package commands

import (
	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var client proto.EasyCarClient

func MustLoad(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connnect easycar :%v", err)
	}
	client = proto.NewEasyCarClient(conn)
}

func GetEasyCarClient() proto.EasyCarClient {
	return client
}
