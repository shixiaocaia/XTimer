package service

import "context"

func (s *XTimerService) ScheduleTask(ctx context.Context) error {
	return s.schedulerUC.Work(ctx)
}
