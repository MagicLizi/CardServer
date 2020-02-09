package fight

import (
	"../config"
	"github.com/gorilla/websocket"
)

var curRooms = make(map[int]*Room)

func CreateFightRoom(creator string, heroId string, pve bool, conn *websocket.Conn) (*Room, bool) {

	room := isInRoom(creator)
	room = nil // always create
	if room != nil {
		room.ResetPlayerConn(creator, conn)
		return room, false
	}

	hero := config.GetHeroById(heroId)
	p1 := InitPlayer(creator, Common, hero, P1, conn)
	p1.InitPlayerLibraries(config.P1Cards)
	newRoom := InitRoom(p1)
	if pve {
		comHero := config.GetHeroById("h1") //pve 写死给 h1 暂时
		newRoom.P2 = InitPlayer("comp2", PC, comHero, P2, nil)
		newRoom.P2.InitPlayerLibraries(config.P2Cards) // pc 暂时写死用2 号卡组
	}
	roomId := newRoom.RoomId
	curRooms[roomId] = newRoom
	return newRoom, true
}

func FindRoom(username string, roomId int) *Room {
	for _, v := range curRooms {
		if v.RoomId == roomId && (v.P1.UserName == username || v.P2.UserName == username) {
			return v
		}
	}
	return nil
}

func isInRoom(creator string) *Room {
	for _, v := range curRooms {
		if v.P1.UserName == creator || v.P2.UserName == creator {
			return v
		}
	}
	return nil
}
