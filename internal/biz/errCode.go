package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrNameNotFound = errors.New(450, "Not Found Name", "没有找到Name")
)
