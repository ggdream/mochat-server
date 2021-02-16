package chat

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)


var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func (r *http.Request) bool {
			return true
		},
	}
	ctx = context.Background()
)
