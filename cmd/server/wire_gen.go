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
	migratorUseCase := biz.NewMigratorUseCase(confData, timerRepo, timerTaskRepo)
	xTimerUseCase := biz.NewXTimerUseCase(confData, timerRepo, transaction, timerTaskRepo, migratorUseCase)
	xTimerService := service.NewXTimerService(xTimerUseCase)
	grpcServer := server.NewGRPCServer(confServer, xTimerService, logger)
	httpServer := server.NewHTTPServer(confServer, xTimerService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}
