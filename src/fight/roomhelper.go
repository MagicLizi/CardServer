package fight

import "time"

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

func CreateFightRoom(creator string) int {

	roomId := isInRoom(creator)

	if roomId == 0 {
		roomId = int(time.Now().Unix())
		room := Room{
			RoomId: roomId,
			P1: Player{
				UserName: creator,
			},
			RState: Ready,
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

func AddComputer(roomId int64) {

}
