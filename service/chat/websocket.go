package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


// Im Websocket通信Control
func Im(c *gin.Context) {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	trans[c.Query("from")] = conn
	defer delete(trans, c.Query("from"))

	handler(conn, c.Query("to"))
}

func handler(conn *websocket.Conn, to string) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		if trans[to] == nil{
			conn.WriteMessage(msgType, []byte("没上线"))
			continue
		}
		if err := trans[to].WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(*(*string)(unsafe.Pointer(&msg)))
	}
}
