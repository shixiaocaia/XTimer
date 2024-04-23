package biz

import (
	"github.com/panjf2000/ants/v2"
	"time"
)

type WorkerPool interface {
	Submit(func()) error
}

type GoWorkerPool struct {
	pool *ants.Pool
}

func (g *GoWorkerPool) Submit(f func()) error {
	return g.pool.Submit(f)
}

func NewGoWorkerPool(size int) *GoWorkerPool {
	pool, err := ants.NewPool(
		size,
		ants.WithExpiryDuration(time.Minute),
	)
	if err != nil {
		panic(err)
	}
	return &GoWorkerPool{pool: pool}
}
