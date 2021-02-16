package chat

import "github.com/gorilla/websocket"


var trans = make(map[string]*websocket.Conn)
