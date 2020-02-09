package config

type Card struct {
	Card_id       int
	Card_name     string
	Card_skin     int
	Card_describe string
	Card_type     int
	Card_hp       int
	Card_cost     int
	Card_combat   int
	Card_tenet    int
	Card_skill1   string
	Card_skill2   string
}

var ConfCards []Card

var Cards map[int]Card = make(map[int]Card)

func GetCardById(cardId int) Card {
	return Cards[cardId]
}

func ParseCards() {
	for _, v := range ConfCards {
		Cards[v.Card_id] = v
	}
}
