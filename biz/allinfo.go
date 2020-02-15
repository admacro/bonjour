package biz

import (
	"github.com/admacro/bonjour/biz/p1"
	"github.com/admacro/bonjour/biz/p2"
)

func AllInfo() string {
	return Info() + p1.Info() + p2.Info()
}
