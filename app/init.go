package app

import (
	"github.com/ggdream/mochat-server/cache"
	"github.com/ggdream/mochat-server/db"
)


//go:generate docker rm -f mo-mysql mo-redis
func initInfra() error {
	if err := db.Init(); err != nil {
		return err
	}
	if err := cache.Init(); err != nil {
		return err
	}
	return nil
}
