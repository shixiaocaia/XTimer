package task

import (
	"Xtimer/internal/conf"
	"Xtimer/internal/service"
	"context"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewTaskServer)

type TaskServer struct {
	service   *service.XTimerService
	scheduler *TaskScheduler
	confData  *conf.Data
}

func (t *TaskServer) Stop(ctx context.Context) error {
	t.scheduler.Stop()
	return nil
}

// 添加Job方法
func (t *TaskServer) NewJobs() []Job {
	return []Job{t.job1, t.job2}
}

// 注入对应service
func NewTaskServer(conf *conf.Server, s *service.XTimerService) *TaskServer {
	t := &TaskServer{
		service: s,
	}
	t.scheduler = NewScheduler(NewTasks(conf.GetTask(), t.NewJobs()))
	return t
}

func NewTasks(c *conf.Server_TASK, jobs []Job) []*Task {
	var tasks []*Task
	for i, job := range jobs {
		tasks = append(tasks, &Task{
			Name:     c.Tasks[i].Name,
			Type:     c.Tasks[i].Type,
			Schedule: c.Tasks[i].Schedule,
			Handler:  job,
		})
	}

	return tasks
}

// 添加Job方法

func (t *TaskServer) job1() {
	t.service.ScheduleTask(context.Background())
}

func (t *TaskServer) job2() {
	t.service.MigratorTimer(context.Background())
}
