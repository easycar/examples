package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuqinqiang/easycar/proto"
	"log"
	"time"
)

type DebitReq struct {
	UserId string `json:"userId"`
	Amount int64  `json:"amount"`
}

type Srv struct {
}

func (srv *Srv) TryDebit(ctx *gin.Context) {
	var (
		req DebitReq
	)
	time.Sleep(300 * time.Millisecond)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, "TryDebit err")
		return
	}
	fmt.Println("[Account]TryDebit req:", req, time.Now().Unix())
	ctx.JSON(200, nil)
}

func (srv *Srv) ConfirmDebit(ctx *gin.Context) {
	var (
		req DebitReq
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, fmt.Sprintf("ConfirmDebit err:%v", err))
		return
	}
	fmt.Println("[Account] ConfirmDebit req:", req)
	ctx.JSON(200, nil)

}

func (srv *Srv) CancelDebit(ctx *gin.Context) {
	var (
		req DebitReq
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(500, "CancelDebit err")
		return
	}
	fmt.Println("[Account] CancelDebit req:", req)
	ctx.JSON(200, nil)

}

func Start(port int) {
	e := gin.Default()
	srv := new(Srv)
	e.POST("/account/tryDebit", srv.TryDebit)
	e.POST("/account/confirmDebit", srv.ConfirmDebit)
	e.POST("/account/cancelDebit", srv.CancelDebit)
	go func() {
		err := e.Run(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("failed the account server:%v", err)
		}
	}()
	fmt.Println("account server start:", port)
}

func NewData() []byte {
	reqData := DebitReq{
		UserId: "remember",
		Amount: 100,
	}
	b, _ := json.Marshal(reqData)
	return b
}

func RegisterTCC(port int) (branches []*proto.RegisterReq_Branch) {
	b := NewData()
	uri := fmt.Sprintf("http://localhost:%d", port)

	// try
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/account/tryDebit",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_TRY,
		Level:    1,
	})
	// confirm
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/account/confirmDebit",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_CONFIRM,
		Level:    1,
	})
	// cancel
	branches = append(branches, &proto.RegisterReq_Branch{
		Uri:      uri + "/account/cancelDebit",
		ReqData:  string(b),
		TranType: proto.TranType_TCC,
		Protocol: "http",
		Action:   proto.Action_CANCEL,
		Level:    1,
	})
	return
}
