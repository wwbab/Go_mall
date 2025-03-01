package service

import (
	"context"

	pbapi "github.com/wwbab/Go_mall/demo/demo_proto/kitex_gen/pbapi"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Respone, err error) {
	// Finish your business logic.

	return &pbapi.Respone{Message: req.Message}, nil
}
