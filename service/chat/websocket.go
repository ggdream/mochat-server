package chat

import (
	"context"
	"fmt"
	"time"
	"unsafe"

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

	cancelCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	go readFromRemote(cancelCtx, conn)
	wirteToRemote(conn)
}

func readFromRemote(ctx context.Context, conn *websocket.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(*(*string)(unsafe.Pointer(&msg)))
		}
	}
}

func wirteToRemote(conn *websocket.Conn) {
	err := conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		fmt.Println(err)
	}
}
