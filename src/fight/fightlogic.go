package fight

import (
	"../config"
	"github.com/gorilla/websocket"
	"log"
)

var curRooms = make(map[int]*Room)

func CreateFightRoom(creator string, heroId string, pve bool, conn *websocket.Conn) (*Room, bool) {

	room := isInRoom(creator)
	if room != nil {
		room.ResetPlayerConn(creator, conn)
		return room, false
	}

	hero := config.GetHeroById(heroId)
	p1 := InitPlayer(creator, Common, hero, P1, conn)
	newRoom := InitRoom(p1)
	if pve {
		comHero := config.GetHeroById("h1") //pve 写死给 h1 暂时
		newRoom.P2 = InitPlayer("comp2", PC, comHero, P2, nil)
	}
	roomId := newRoom.RoomId
	curRooms[roomId] = newRoom
	return newRoom, true
}

func isInRoom(creator string) *Room {
	for _, v := range curRooms {
		if v.P1.UserName == creator || v.P2.UserName == creator {
			return v
		}
	}
	return nil
}

func simpleFindRoom(roomId int) *Room {
	for _, v := range curRooms {
		if v.RoomId == roomId {
			return v
		}
	}
	return nil
}

func findRoom(username string, roomId int) *Room {
	for _, v := range curRooms {
		if v.RoomId == roomId && (v.P1.UserName == username || v.P2.UserName == username) {
			return v
		}
	}
	return nil
}

func PlayerRoomReady(username string, roomId int) uint32 {
	room := findRoom(username, roomId)
	if room == nil {
		return 0
	}
	player := room.GetPlayer(username)
	if room.State == Waiting {
		//如果两人都ready了就开战
		player.State = Fight
		room.TryStartFight()
	} else if room.State == Fighting {
		log.Println("断线重连")
		room.PlayerReConnect(player)
	}
	return 1
}
