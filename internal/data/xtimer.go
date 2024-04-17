package data

import (
	"context"

	"Xtimer/internal/biz"
)

type xTimerRepo struct {
	data *Data
}

func NewXTimerRepo(data *Data) biz.XTimerRepo {
	return &xTimerRepo{
		data: data,
	}
}

func (r *xTimerRepo) Save(ctx context.Context, g *biz.Timer) (*biz.Timer, error) {
	// 开启事务的话, 需要调用r.data.DB(ctx) 而不是r.data.db
	err := r.data.DB(ctx).Create(g).Error
	return g, err
}

func (r *xTimerRepo) Update(ctx context.Context, g *biz.Timer) (*biz.Timer, error) {
	err := r.data.db.WithContext(ctx).Where("id = ?", g.TimerId).Updates(g).Error
	return g, err
}

func (r *xTimerRepo) Delete(ctx context.Context, id int64) error {
	return r.data.DB(ctx).Where("id = ?", id).Delete(&biz.Timer{}).Error
}

func (r *xTimerRepo) FindByID(ctx context.Context, timerId int64) (*biz.Timer, error) {
	var timer biz.Timer
	err := r.data.db.WithContext(ctx).Where("id = ?", timerId).First(&timer).Error
	if err != nil {
		return nil, err
	}
	return &timer, nil
}

func (r *xTimerRepo) FindByStatus(ctx context.Context, status int) ([]*biz.Timer, error) {
	var timers []*biz.Timer
	err := r.data.db.WithContext(ctx).Where("status = ?", status).Find(&timers).Error
	if err != nil {
		return nil, err
	}
	return timers, nil
}
