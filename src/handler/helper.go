package handler

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
)

var handlerMap = make(map[uint16]func(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn))

//这个后期需要自动生成
func InitHandlerConf() {
	handlerMap[1] = TempLogin
	handlerMap[2] = CreateRoom
	handlerMap[3] = JoinRoom
	handlerMap[4] = RoomReady
}

func GetHandlerByMsgId(msgId uint16) func(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {
	return handlerMap[msgId]
}

func ParseProto(protoStruct proto.Message, protoBufData []byte) {
	protoParseError := proto.Unmarshal(protoBufData, protoStruct)
	if protoParseError != nil {
		log.Println("parse error", protoParseError)
	}
}

func ToProtoRes(protoStruct proto.Message, msgTypeBuf []byte, conn *websocket.Conn) {
	if data, err := proto.Marshal(protoStruct); err == nil {
		lastData := append(msgTypeBuf, data...)
		log.Println("lastData", lastData)
		conn.WriteMessage(websocket.BinaryMessage, lastData)
	} else {
		log.Println("to proto error", err)
	}
}
