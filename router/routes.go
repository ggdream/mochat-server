package router

import (
	"github.com/ggdream/mochat-server/service/chat"
	"github.com/ggdream/mochat-server/service/sign"
	"github.com/gin-gonic/gin"
)


// SetRoutes 给应用app设置路由（路由出口）
func SetRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/login", sign.Login)
		user.GET("/register", sign.Register)
	}

	im := r.Group("/chat")
	{
		im.GET("/msg", chat.GetMsg)
	}
}
