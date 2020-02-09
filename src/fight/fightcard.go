package fight

import "../config"

type FCard struct {
	Cid        string
	StaticData config.Card
}

func InitFCard(fcId string, cardId int) FCard {
	return FCard{
		Cid:        fcId,
		StaticData: config.GetCardById(cardId),
	}
}
