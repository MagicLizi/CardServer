package handler

import (
	"../protos/server"
	"github.com/gorilla/websocket"
)

func TempLogin(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {
	var target protos.TempLoginReq
	ParseProto(&target, protoBufData)

	ToProtoRes(&protos.TempLoginRes{
		Status:   "1",
		Username: target.GetName(),
	}, msgTypeBuf, conn)
}
