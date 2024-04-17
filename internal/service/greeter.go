package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "server/api/x_timer/v1"
	"server/internal/biz"
)

type XTimerService struct {
	v1.UnimplementedXTimerServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *XTimerService {
	return &XTimerService{uc: uc}
}

// SayHello implements x_timer.GreeterServer.
func (s *XTimerService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	log.Infof("SayHello: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, biz.ErrNameNotFound
	}

	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
