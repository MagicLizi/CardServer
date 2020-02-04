package handler

import (
	"../fight"
	"../protos/server"
	"github.com/gorilla/websocket"
	"log"
)

func CreateRoom(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {
	log.Println("try create")
	var target protos.CreateRoomReq
	ParseProto(&target, protoBufData)

	ToProtoRes(&protos.CreateRoomRes{
		RoomId: uint32(fight.CreateFightRoom(target.Username, target.HeroId, target.Pve, conn)),
	}, msgTypeBuf, conn)
}

func JoinRoom(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {
	log.Println("join create")
}

func RoomReady(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {
	log.Println("RoomReady")
	var target protos.RoomReadyReq
	ParseProto(&target, protoBufData)

	ToProtoRes(&protos.RoomReadyRes{
		Result: fight.PlayerRoomReady(target.Username, int(target.RoomId)),
	}, msgTypeBuf, conn)
}
