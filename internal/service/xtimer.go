package service

import (
	v1 "Xtimer/api/x_timer/v1"
	"Xtimer/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/goccy/go-json"
)

type XTimerService struct {
	v1.UnimplementedXTimerServer

	timerUC *biz.XTimerUseCase
}

func NewXTimerService(uc *biz.XTimerUseCase) *XTimerService {
	return &XTimerService{timerUC: uc}
}

// SayHello implements x_timer.GreeterServer.
func (s *XTimerService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	log.Infof("SayHello: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, biz.ErrNameNotFound
	}

	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *XTimerService) CreateTimer(ctx context.Context, req *v1.CreateTimerRequest) (*v1.CreateTimerReply, error) {
	param, err := json.Marshal(req.NotifyHTTPParam)
	if err != nil {
		return nil, err
	}
	timer, err := s.timerUC.CreateTimer(ctx, &biz.Timer{
		App:             req.App,
		Name:            req.Name,
		Status:          0,
		Cron:            req.Cron,
		NotifyHTTPParam: string(param),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateTimerReply{
		Id: int32(timer.TimerId),
	}, nil
}

func (s *XTimerService) ActiveTimer(ctx context.Context, req *v1.ActiveTimerRequest) (*v1.ActiveTimerReply, error) {
	err := s.timerUC.ActiveTimer(ctx, req.App, req.Id, req.Status)
	if err != nil {
		return nil, err
	}
	return &v1.ActiveTimerReply{
		Id:      req.Id,
		Message: "Active timer success",
	}, nil
}
