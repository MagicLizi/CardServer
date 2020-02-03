package fight

import (
	"../config"
	"github.com/gorilla/websocket"
	"time"
)

type RoomState int

const (
	Ready    RoomState = 1
	Fighting RoomState = 2
)

type Room struct {
	RoomId int
	P1     Player
	P2     Player
	RState RoomState
}

var curRooms = make(map[int]Room)

func CreateFightRoom(creator string, heroId string, pve bool, conn *websocket.Conn) int {

	roomId := isInRoom(creator)

	if roomId == 0 {
		roomId = int(time.Now().Unix())
		hero := config.GetHeroById(heroId)
		room := Room{
			RoomId: roomId,
			P1: Player{
				UserName: creator,
				Hero: FHero{
					Name:       hero.Hero_name,
					CurHp:      hero.Hero_hp,
					StaticData: *hero,
				},
				Conn: conn,
				Type: Common,
			},
			RState: Ready,
		}
		if pve {
			comHero := config.GetHeroById("h1") //pve 写死给 h1 暂时
			room.P2 = Player{
				UserName: "ComP2",
				Hero: FHero{
					Name:       comHero.Hero_name,
					CurHp:      comHero.Hero_hp,
					StaticData: *comHero,
				},
				Type: PC,
			}
		}

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
