package data

import (
	"Xtimer/internal/conf"
	"Xtimer/third_party/gormcli"
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewXTimerRepo, NewDatabase)

type contextTxKey struct{}

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(db *gorm.DB) *Data {
	dt := &Data{db: db}
	return dt
}

func NewDatabase(conf *conf.Data) *gorm.DB {
	dt := conf.GetDatabase()
	gormcli.Init(
		gormcli.WithAddr(dt.GetAddr()),
		gormcli.WithUser(dt.GetUser()),
		gormcli.WithPassword(dt.GetPassword()),
		gormcli.WithDataBase(dt.GetDatabase()),
		gormcli.WithMaxIdleConn(int(dt.GetMaxIdleConn())),
		gormcli.WithMaxOpenConn(int(dt.GetMaxOpenConn())),
		gormcli.WithMaxIdleTime(int64(dt.GetMaxIdleTime())),
		// 如果设置了慢日志阈值, 则打印日志，否则不打印
		gormcli.WithSlowThresholdMillisecond(dt.GetSlowThresholdMillisecond()))

	return gormcli.GetDB()
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}
