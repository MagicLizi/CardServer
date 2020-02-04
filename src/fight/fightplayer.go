package fight

import (
	"../config"
	"github.com/gorilla/websocket"
)

type PlayerType int

const (
	Common PlayerType = 1
	PC     PlayerType = 2
)

type PlayerState int

const (
	Ready PlayerState = 1
	Fight PlayerState = 2
	Exit  PlayerState = 3
)

type Player struct {
	UserName     string
	Hero         FHero
	Conn         *websocket.Conn
	Type         PlayerType
	State        PlayerState
	Library      map[string]FCard //当前牌库
	DisLibrary   map[string]FCard //当前弃牌堆
	HandLibrary  map[string]FCard //当前手牌
	RoundLibrary map[string]FCard //当前回合打出得牌堆
}

func InitPlayer(username string, t PlayerType, hero *config.Hero, conn *websocket.Conn) Player {
	p := Player{
		UserName: username,
		Hero: FHero{
			Name:       hero.Hero_name,
			CurHp:      hero.Hero_hp,
			StaticData: *hero,
		},
		Conn:         conn,
		Type:         t,
		State:        Ready,
		Library:      map[string]FCard{},
		DisLibrary:   map[string]FCard{},
		HandLibrary:  map[string]FCard{},
		RoundLibrary: map[string]FCard{},
	}
	if t == PC {
		p.State = Fight
	}
	return p
}
