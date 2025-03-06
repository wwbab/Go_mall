package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	api "github.com/wwbab/Go_mall/demo/demo_thrift/kitex_gen/api"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(request *api.Request) (resp *api.Response, err error) {
	// Finish your business logic.
	info := rpcinfo.GetRPCInfo(s.ctx)
	fmt.Println(info.From().ServiceName())
	
	return &api.Response{Message: request.Message},nil
}
