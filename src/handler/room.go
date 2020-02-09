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

	room, create := fight.CreateFightRoom(target.Username, int(target.HeroId), target.Pve, conn)

	if target.Pve {
		ToProtoRes(&protos.CreateRoomRes{
			RoomId: uint32(room.RoomId),
			P1: &protos.PlayerInfo{
				Username: room.P1.UserName,
				HeroName: room.P1.Hero.Name,
				HeroHp:   uint32(room.P1.Hero.CurHp),
				Fv:       uint32(room.P1.Hero.CurFV),
				Belief:   uint32(room.P1.Hero.CurBelief),
			},
			P2: &protos.PlayerInfo{
				Username: room.P2.UserName,
				HeroName: room.P2.Hero.Name,
				HeroHp:   uint32(room.P2.Hero.CurHp),
				Fv:       uint32(room.P2.Hero.CurFV),
				Belief:   uint32(room.P2.Hero.CurBelief),
			},
			IsCreate: create,
		}, msgTypeBuf, conn)
	} else {
		if create && &room.P2 == nil {
			ToProtoRes(&protos.CreateRoomRes{
				RoomId: uint32(room.RoomId),
				P1: &protos.PlayerInfo{
					Username: room.P1.UserName,
					HeroName: room.P1.Hero.Name,
					HeroHp:   uint32(room.P1.Hero.CurHp),
					Fv:       uint32(room.P1.Hero.CurFV),
					Belief:   uint32(room.P1.Hero.CurBelief),
				},
				IsCreate: create,
			}, msgTypeBuf, conn)
		} else if !create && &room.P2 != nil {
			ToProtoRes(&protos.CreateRoomRes{
				RoomId: uint32(room.RoomId),
				P1: &protos.PlayerInfo{
					Username: room.P1.UserName,
					HeroName: room.P1.Hero.Name,
					HeroHp:   uint32(room.P1.Hero.CurHp),
					Fv:       uint32(room.P1.Hero.CurFV),
					Belief:   uint32(room.P1.Hero.CurBelief),
				},
				P2: &protos.PlayerInfo{
					Username: room.P2.UserName,
					HeroName: room.P2.Hero.Name,
					HeroHp:   uint32(room.P2.Hero.CurHp),
					Fv:       uint32(room.P2.Hero.CurFV),
					Belief:   uint32(room.P2.Hero.CurBelief),
				},
				IsCreate: create,
			}, msgTypeBuf, conn)
		}
	}
}

func RoomReady(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {
	log.Println("RoomReady")
	var target protos.RoomReadyReq
	ParseProto(&target, protoBufData)

	roomId := int(target.RoomIde.RoomId)
	username := target.RoomIde.Username
	room := fight.FindRoom(username, roomId)
	if room == nil {
		log.Println("room not exist")
		return
	}
	room.RoomReady(username)
	//ToProtoRes(&protos.RoomReadyRes{
	//	Result: room.RoomReady(username),
	//}, msgTypeBuf, conn)
}

func RenderCenterShopEnd(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {
	log.Println("render end")
	var target protos.RenderCenterShopEnd
	ParseProto(&target, protoBufData)

	roomId := int(target.RoomIde.RoomId)
	username := target.RoomIde.Username
	room := fight.FindRoom(username, roomId)
	if room == nil {
		log.Println("room not exist")
		return
	}
	room.CenterShopRenderEnd(username)
}

func RenderLotteryHandCardsEnd(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {

}

func JoinRoom(protoBufData []byte, msgTypeBuf []byte, conn *websocket.Conn) {
	log.Println("join create")
}
