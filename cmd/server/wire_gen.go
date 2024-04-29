// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"Xtimer/internal/biz"
	"Xtimer/internal/conf"
	"Xtimer/internal/data"
	"Xtimer/internal/server"
	"Xtimer/internal/service"
	"Xtimer/internal/task"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDatabase(confData)
	client := data.NewCache(confData)
	dataData := data.NewData(db, client)
	timerRepo := data.NewTimerRepo(dataData)
	transaction := data.NewTransaction(dataData)
	timerTaskRepo := data.NewTimerTaskRepo(dataData)
	taskCacheRepo := data.NewTaskCacheRepo(confData, dataData)
	httpClient := biz.NewHttpClient()
	migratorUseCase := biz.NewMigratorUseCase(confData, timerRepo, timerTaskRepo, taskCacheRepo)
	xTimerUseCase := biz.NewXTimerUseCase(confData, timerRepo, transaction, timerTaskRepo, migratorUseCase, taskCacheRepo)
	executorUseCase := biz.NewExecutorUseCase(confData, timerRepo, timerTaskRepo, httpClient)
	triggerUseCase := biz.NewTriggerUseCase(confData, timerRepo, timerTaskRepo, transaction, taskCacheRepo, executorUseCase)
	schedulerUseCase := biz.NewSchedulerUseCase(confData, timerRepo, timerTaskRepo, taskCacheRepo, transaction, triggerUseCase)
	xTimerService := service.NewXTimerService(xTimerUseCase, schedulerUseCase, migratorUseCase)
	grpcServer := server.NewGRPCServer(confServer, xTimerService, logger)
	httpServer := server.NewHTTPServer(confServer, xTimerService, logger)
	taskServer := task.NewTaskServer(confServer, xTimerService)
	app := newApp(logger, grpcServer, httpServer, taskServer)
	return app, func() {
	}, nil
}
