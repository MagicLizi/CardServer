package fight

import "github.com/gorilla/websocket"

type PlayerType int

const (
	Common PlayerType = 1
	PC     PlayerType = 2
)

type Player struct {
	UserName string
	Hero     FHero
	Conn     *websocket.Conn
	Type     PlayerType
}
