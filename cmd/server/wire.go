//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

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
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, task.ProviderSet, newApp))
}
