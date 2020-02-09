package fight

import (
	"../config"
	"../notify"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"strings"
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

type Side int

const (
	P1 Side = 1
	P2 Side = 2
)

type RenderState int

const (
	Rendering RenderState = 1
	RenderEnd RenderState = 2
)

type Player struct {
	UserName              string
	Hero                  FHero
	Conn                  *websocket.Conn
	Type                  PlayerType
	State                 PlayerState
	Library               map[string]FCard //当前牌库
	DisLibrary            map[string]FCard //当前弃牌堆
	HandLibrary           map[string]FCard //当前手牌
	RoundLibrary          map[string]FCard //当前回合打出得牌堆
	BuildLibrary          map[string]FCard //当前构筑牌堆
	PSide                 Side
	CenterShopRenderState RenderState
}

func InitPlayer(username string, t PlayerType, hero *config.Hero, s Side, conn *websocket.Conn) *Player {
	p := Player{
		UserName: username,
		Hero: FHero{
			Name:       hero.Hero_name,
			CurHp:      hero.Hero_hp,
			StaticData: *hero,
		},
		Conn:                  conn,
		Type:                  t,
		State:                 Ready,
		Library:               map[string]FCard{}, //当前待抽取的牌库
		DisLibrary:            map[string]FCard{}, //废弃牌库
		HandLibrary:           map[string]FCard{}, //当前手牌
		RoundLibrary:          map[string]FCard{}, //每回合使用过的牌的牌库
		BuildLibrary:          map[string]FCard{}, //带入战斗的构筑牌库
		PSide:                 s,
		CenterShopRenderState: RenderEnd,
	}
	if t == PC {
		p.State = Fight
	}
	return &p
}

func (p *Player) Notify(protoStruct proto.Message, messageKey string) {
	if p.Conn != nil && p.Type == Common {
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, notify.GetNotifyIdWithKey(messageKey))
		notify.ToProtoNotify(protoStruct, b, p.Conn)
	}
}

func (p *Player) AddHandCard(fCardId string, cardId int) {
	p.HandLibrary[fCardId] = InitFCard(fCardId, cardId)
}

func (p *Player) InitPlayerLibraries(buildCard []config.PlayerCard) {
	//英雄卡组 加入 Library
	hCards := strings.Split(p.Hero.StaticData.Hero_card, ",")
	for _, v := range hCards {
		cardInfo := strings.Split(v, "_")
		cardId, err := strconv.Atoi(cardInfo[0])
		count, err := strconv.Atoi(cardInfo[1])
		if err == nil {
			for i := 0; i < count; i++ {
				fCardId := fmt.Sprintf("%d_%d_%d", cardId, i, p.PSide)
				fCard := InitFCard(fCardId, cardId)
				p.Library[fCardId] = fCard
			}
		}
	}

	//初始化构筑卡组
	for _, v := range buildCard {
		count := v.Card_number
		cardId, err := strconv.Atoi(v.Card_id)
		if err != nil {
			for i := 0; i < count; i++ {
				fCardId := fmt.Sprintf("%d_%d_%d", cardId, i, p.PSide)
				fCard := InitFCard(fCardId, cardId)
				p.BuildLibrary[fCardId] = fCard
			}
		}
	}
	log.Println("InitPlayerLibraries success!")
}

func (p *Player) LotteryCardsToHand() {
	//计算当次抽取几张
	handCardsLimit := 10
	onceCards := 5
	leftCards := handCardsLimit - len(p.HandLibrary)
	needLotteryCount := 0
	if leftCards >= onceCards {
		needLotteryCount = onceCards
	} else {
		needLotteryCount = leftCards
	}

	var ids []string
	//计算是否需要 洗牌
	if len(p.Library) >= needLotteryCount {
		ids = p.LotteryFromLibrary(needLotteryCount)
	} else {
		for k, v := range p.Library {
			ids = append(ids, v.Cid)
			delete(p.Library, k)
		}

		p.ResetLibrary()

		leftNeed := needLotteryCount - len(ids)
		leftIds := p.LotteryFromLibrary(leftNeed)
		ids = append(ids, leftIds...)
	}
}

func (p *Player) LotteryFromLibrary(needLotteryCount int) []string {
	var randomIds []FCard
	for _, v := range p.Library {
		randomIds = append(randomIds, v)
	}

	var ids []string
	result := LotteryCards(randomIds, needLotteryCount)
	for _, v := range result {
		p.HandLibrary[v.Cid] = v
		delete(p.Library, v.Cid)
		ids = append(ids, v.Cid)
	}

	return ids
}

func (p *Player) ResetLibrary() {

	for k, _ := range p.Library {
		delete(p.Library, k)
	}

	for k, v := range p.DisLibrary {
		p.Library[k] = v
		delete(p.DisLibrary, k)
	}
}

func (p *Player) AddCardToDisLibrary(card FCard) {
	p.DisLibrary[card.Cid] = card
}
