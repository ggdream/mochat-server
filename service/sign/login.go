package sign

import (
	"crypto/md5"
	"fmt"
	"github.com/ggdream/mochat-server/cache"
	"github.com/ggdream/mochat-server/db"
	"github.com/ggdream/mochat-server/model/user"
	"github.com/gin-gonic/gin"
	"unsafe"
)


// Login 用户登录
func Login(c *gin.Context) {
	ac := user.NewAccount()
	if err := c.ShouldBind(ac); err != nil {
		c.JSON(200, gin.H{"code": -1, "data": "参数错误"})
		return
	}

	if err := db.SearchUser(ac); err != nil {
		c.JSON(200, gin.H{"code": -2, "data": "登录失败"})
		return
	}

	token := fmt.Sprintf("%x", md5.Sum(*(*[]byte)(unsafe.Pointer(&ac.Password))))
	if err := cache.Login(ac.Username, token); err != nil {
		c.JSON(200, gin.H{"code": -3, "data": "登录失败"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "data": token})
}
