package service

import "context"

func (s *XTimerService) MigratorTimer(ctx context.Context) error {
	return s.migratorUC.BatchMigratorTimer(ctx)
}
