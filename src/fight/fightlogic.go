package fight

import (
	"../config"
	"github.com/gorilla/websocket"
)

var curRooms = make(map[int]Room)

func CreateFightRoom(creator string, heroId string, pve bool, conn *websocket.Conn) int {

	roomId := isInRoom(creator)

	if roomId == 0 {

		hero := config.GetHeroById(heroId)
		p1 := InitPlayer(creator, Common, hero, conn)
		room := InitRoom(p1)

		if pve {
			comHero := config.GetHeroById("h1") //pve 写死给 h1 暂时
			room.P2 = InitPlayer("comp2", PC, comHero, nil)
		}
		roomId = room.RoomId
		curRooms[roomId] = room
	}

	return roomId
}

func isInRoom(creator string) int {
	for _, v := range curRooms {
		if v.P1.UserName == creator || v.P2.UserName == creator {
			return v.RoomId
		}
	}
	return 0
}

func simpleFindRoom(roomId int) *Room {
	for _, v := range curRooms {
		if v.RoomId == roomId {
			return &v
		}
	}
	return nil
}

func findRoom(username string, roomId int) *Room {
	for _, v := range curRooms {
		if v.RoomId == roomId && (v.P1.UserName == username || v.P2.UserName == username) {
			return &v
		}
	}
	return nil
}

func PlayerRoomReady(username string, roomId int) uint32 {
	room := findRoom(username, roomId)
	if room == nil {
		return 0
	}
	//如果两人都ready了就开战
	player := room.GetPlayer(username)
	player.State = Fight
	room.TryStartFight()
	return 1
}
