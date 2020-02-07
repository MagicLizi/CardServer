package fight

import (
	"../config"
	"../protos/server"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"time"
)

type RoomState int

const (
	Waiting  RoomState = 1
	Fighting RoomState = 2
	End      RoomState = 3
)

type Room struct {
	RoomId        int
	P1            *Player
	P2            *Player
	CenterLibrary map[string]FCard //中央牌库
	CenterShop    map[string]FCard //中央购买区
	CurTurn       *Player          //当前行动方
	State         RoomState
}

func (r *Room) ResetPlayerConn(username string, conn *websocket.Conn) {
	log.Println("ResetPlayerConn")
	player := r.GetPlayer(username)
	if player != nil {
		player.Conn = conn
	}
}

func (r *Room) ChangeTurn() {
	luckCardId := "c0"
	if r.CurTurn == nil {
		turn := rand.Intn(2)
		if turn == 0 {
			r.CurTurn = r.P1
			r.P2.AddHandCard(fmt.Sprintf("%s_%d_%d", luckCardId, 0, r.P2.PSide), luckCardId)

		} else {
			r.CurTurn = r.P2
			r.P1.AddHandCard(fmt.Sprintf("%s_%d_%d", luckCardId, 0, r.P1.PSide), luckCardId)
		}
	} else {
		if r.CurTurn == r.P1 {
			r.CurTurn = r.P2
		} else {
			r.CurTurn = r.P1
		}
	}
}

func (r *Room) InitCenterLibrary() {

	//玩家1 和玩家2 的各20张

	for k, v := range r.P1.BuildLibrary {
		r.CenterLibrary[k] = v
	}

	for k, v := range r.P2.BuildLibrary {
		r.CenterLibrary[k] = v
	}

	//中央牌库 抽取20张 因为存在必然存在的，所以先将必然存在的加入中央牌库
	mustCount := 0
	for _, v := range config.MustCenterCards {
		cardId := v.Card_id
		count := v.Card_number
		for i := 0; i < count; i++ {
			fCardId := fmt.Sprintf("%s_%d_center", cardId, i)
			r.CenterLibrary[fCardId] = InitFCard(fCardId, cardId)
			mustCount++
		}
	}

	leftNeedCount := 20 - mustCount
	var randomIds []FCard
	for _, v := range config.RandomCenterCards {
		cardId := v.Card_id
		count := v.Card_number
		for i := 0; i < count; i++ {
			fCardId := fmt.Sprintf("%s_%d_center", cardId, i)
			randomIds = append(randomIds, InitFCard(fCardId, cardId))
		}
	}

	randomFCards := lotteryCards(randomIds, leftNeedCount)
	for _, v := range randomFCards {
		r.CenterLibrary[v.Cid] = v
	}

	log.Println("Init CenterLibrary Success")
}

func (r *Room) GetPlayer(username string) *Player {
	if r.P1.UserName == username {
		return r.P1
	}
	return r.P2
}

func (r *Room) TryStartFight() {
	if r.P1 != nil && r.P1.State == Fight && r.P2 != nil && r.P2.State == Fight {
		log.Println("all player is ready begin fight!")
		r.P1.InitPlayerLibraries(config.P1Cards)
		r.P2.InitPlayerLibraries(config.P2Cards)
		r.InitCenterLibrary()
		r.RefreshCenterShop()
		r.State = Fighting
	}
}

func (r *Room) RefreshCenterShop() {
	centerShopCount := 6
	needLotteryCount := centerShopCount - len(r.CenterShop)
	var refreshCards []string

	//获取所有CenterLibrary 中得卡片数组
	var randomIds []FCard
	for _, v := range r.CenterLibrary {
		randomIds = append(randomIds, v)
	}

	result := lotteryCards(randomIds, needLotteryCount)

	for _, v := range result {
		refreshCards = append(refreshCards, v.Cid)
		r.CenterShop[v.Cid] = v
		delete(r.CenterLibrary, v.Cid)
	}

	log.Println("refresh shop end")
	//通知客户端CenterShop刷新

	r.P1.Notify(&protos.NotifyRefreshCenterShop{
		RefreshCards: refreshCards,
	}, "RefreshCenterShop")

}

func (r *Room) PlayerReConnect(player *Player) {

	var centerShop []string
	var handCards []string

	for k, _ := range r.CenterShop {
		centerShop = append(centerShop, k)
	}

	for k, _ := range player.HandLibrary {
		handCards = append(handCards, k)
	}

	player.Notify(&protos.NotifyPlayerRoomInfo{
		CenterShopCards: centerShop,
		PlayerHandCards: handCards,
	}, "PlayerRoomInfo")
}

func InitRoom(p *Player) *Room {
	roomId := int(time.Now().Unix())
	return &Room{
		RoomId:        roomId,
		CenterLibrary: map[string]FCard{},
		CenterShop:    map[string]FCard{},
		P1:            p,
		State:         Waiting,
	}
}

func lotteryCards(from []FCard, count int) []FCard {
	var result []FCard
	for i := 0; i < count; i++ {
		lotteryIndex := rand.Intn(len(from))
		lotteryCard := from[lotteryIndex]
		from = append(from[:lotteryIndex], from[lotteryIndex+1:]...)
		result = append(result, lotteryCard)
	}
	return result
}
