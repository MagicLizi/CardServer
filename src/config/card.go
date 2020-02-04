package config

type Card struct {
	Card_id       string
	Card_name     string
	Card_skin     string
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

var Cards map[string]Card = make(map[string]Card)

func GetCardById(cardId string) Card {
	return Cards[cardId]
}

func ParseCards() {
	for _, v := range ConfCards {
		Cards[v.Card_id] = v
	}
}
