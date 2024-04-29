package biz

import (
	v1 "Xtimer/api/x_timer/v1"
	"Xtimer/internal/conf"
	"Xtimer/internal/constant"
	"Xtimer/internal/utils"
	"fmt"
	"time"

	"Xtimer/third_party/log"
	"context"
	"encoding/json"
)

type ExecutorUseCase struct {
	confData   *conf.Data
	httpClient *JSONClient
	timerRepo  TimerRepo
	taskRepo   TimerTaskRepo
}

func NewExecutorUseCase(confData *conf.Data, timerRepo TimerRepo, taskRepo TimerTaskRepo, httpClient *JSONClient) *ExecutorUseCase {
	return &ExecutorUseCase{
		confData:   confData,
		timerRepo:  timerRepo,
		taskRepo:   taskRepo,
		httpClient: httpClient,
	}
}

func (w *ExecutorUseCase) Work(ctx context.Context, timerIDUnixKey string) error {
	// 拿到消息，查询一次完整的 timer 定义
	timerID, unix, err := utils.SplitTimerIDUnix(timerIDUnixKey)
	if err != nil {
		return err
	}

	// 幂等性验证
	// todo 1. bloomFilter
	// 2. mysql
	task, err := w.taskRepo.GetTasksByTimerIdAndRunTimer(ctx, timerID, unix)
	if err != nil {
		return fmt.Errorf("get task failed, timerID: %d, runTimer: %d, err: %w", timerID, time.UnixMilli(unix), err)
	}
	if task.Status != constant.TaskStatusNotRun.ToInt() {
		log.WarnContextf(ctx, "task is already executed, timerID: %d, exec_time: %v", timerID, task.RunTimer)
		return nil
	}

	return w.executeAndPostProcess(ctx, timerID, unix)
}

func (w *ExecutorUseCase) executeAndPostProcess(ctx context.Context, timerID int64, unix int64) error {
	// 未执行，则查询 timer 完整的定义，执行回调
	timer, err := w.timerRepo.FindByID(ctx, timerID)
	if err != nil {
		return fmt.Errorf("get timer failed, id: %d, err: %w", timerID, err)
	}

	// 定时器已经处于去激活态，则无需处理任务
	if timer.Status != constant.Enabled.ToInt() {
		log.WarnContextf(ctx, "timer has alread been unabled, timerID: %d", timerID)
		return nil
	}

	execTime := time.Now()
	resp, err := w.execute(ctx, timer)
	return w.postProcess(ctx, resp, err, timer.App, uint(timerID), unix, execTime)
}

func (w *ExecutorUseCase) execute(ctx context.Context, timer *Timer) (map[string]interface{}, error) {
	var (
		// resp map[string]interface{}
		err error
	)
	notifyHTTPParam := v1.NotifyHTTPParam{}
	err = json.Unmarshal([]byte(timer.NotifyHTTPParam), &notifyHTTPParam)
	if err != nil {
		log.Errorf("json unmarshal for NotifyHTTPParam err %s", err.Error())
		return nil, err
	}

	// 暂时支持post
	// err = w.httpClient.Post(ctx, notifyHTTPParam.Url, notifyHTTPParam.Headers, notifyHTTPParam.Body, &resp)
	// return resp, err
	time.Sleep(100 * time.Millisecond)
	return map[string]interface{}{
		"code":   200,
		"status": "success",
		"msg":    "hello world",
	}, nil
}

func (w *ExecutorUseCase) postProcess(ctx context.Context, resp map[string]interface{}, execErr error, app string, timerID uint, unix int64, execTime time.Time) error {
	// todo 上报监控
	// todo 更新bloomFilter

	task, err := w.taskRepo.GetTasksByTimerIdAndRunTimer(ctx, int64(timerID), unix)
	if err != nil {
		return fmt.Errorf("get task failed, timerID: %d, runTimer: %d, err: %w", timerID, time.UnixMilli(unix), err)
	}

	// output
	if execErr != nil {
		task.Output = execErr.Error()
	} else {
		respBody, _ := json.Marshal(resp)
		task.Output = string(respBody)
	}

	// Status
	if execErr != nil {
		task.Status = constant.TaskStatusFail.ToInt()
	} else {
		task.Status = constant.TaskStatusSuccess.ToInt()
	}

	// costTime
	task.CostTime = int(execTime.UnixMilli() - task.RunTimer)

	_, err = w.taskRepo.Update(ctx, task)
	if err != nil {
		return fmt.Errorf("task postProcess failed, timerID: %d, runTimer: %d, err: %w", timerID, time.UnixMilli(unix), err)
	}

	log.InfoContextf(ctx, "task postProcess success, timerID: %d, runTimer: %d", timerID, time.UnixMilli(unix))
	return nil
}
