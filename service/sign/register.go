package sign

import (
	"database/sql"
	"github.com/ggdream/mochat-server/db"
	"github.com/ggdream/mochat-server/model/user"
	"github.com/gin-gonic/gin"
)


// Register 用户注册
func Register(c *gin.Context) {
	ac := user.NewAccount()
	if err := c.Bind(ac); err != nil {
		c.JSON(200, gin.H{"code": -1, "data": "参数错误"})
		return
	}

	if err := db.SearchUser(ac); err == nil || err != sql.ErrNoRows {
		c.JSON(200, gin.H{"code": -2, "data": "注册失败"})
		return
	}

	if err := db.AddUser(ac); err != nil {
		c.JSON(200, gin.H{"code": -3, "data": "注册失败"})
		return
	}

	c.JSON(200, gin.H{"code": 0, "data": "注册成功"})
}
