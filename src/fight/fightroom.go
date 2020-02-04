package fight

import (
	"log"
	"time"
)

type Room struct {
	RoomId        int
	P1            Player
	P2            Player
	CenterLibrary map[string]FCard //中央牌库
	CenterShop    map[string]FCard //中央购买区
}

//初始化 房间循环
//1 初始化中央牌库 P1_20 + P2_20 + 中央牌库20
//2 选择先手方
func (r *Room) InitFightLibraries() {
	log.Println("Init Fight Libraries")
	//玩家1 和玩家2 得 各20张

	//中央牌库 抽取20 张
	//for _, v := range config.ConfCenterCards {
	//
	//}
}

func (r *Room) GetPlayer(username string) *Player {
	if r.P1.UserName == username {
		return &r.P1
	}
	return &r.P2
}

func (r *Room) TryStartFight() {
	if r.P1.State == Fight && r.P2.State == Fight {
		log.Println("all player is ready begin fight!")
		r.InitFightLibraries()
	}
}

func InitRoom(p Player) Room {
	roomId := int(time.Now().Unix())
	return Room{
		RoomId:        roomId,
		CenterLibrary: map[string]FCard{},
		CenterShop:    map[string]FCard{},
		P1:            p,
	}
}
