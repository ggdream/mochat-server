package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ggdream/mochat-server/router"
)


// New 新建应用
func New() error {
	if err := initInfra(); err != nil {
		return err
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router.SetRoutes(r)

	return r.Run()
}
