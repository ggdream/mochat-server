package chat

import (
	"github.com/ggdream/mochat-server/model/msg"
	"github.com/gin-gonic/gin"
)


// GetMsg 获取历史聊天记录
func GetMsg(c *gin.Context) {
	gm := msg.NewDescMsg()
	if err := c.ShouldBind(gm); err != nil {
		c.JSON(200, gin.H{"code": -1, "data": "参数错误"})
		return
	}

	//res, err := 1, erro
	//if err != nil {
	//	c.JSON(200, gin.H{"code": -2, "data": "查询失败"})
	//	return
	//}

	c.JSON(200, gin.H{"code": 0, "data": 1})
}
