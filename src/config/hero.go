package config

type Hero struct {
	Hero_id       string
	Hero_name     string
	Hero_skin     string
	Hero_describe string
	Hero_hp       int
	Hero_card1    string
	Hero_card2    string
	Hero_card3    string
	Hero_card4    string
	Hero_card5    string
	Hero_card6    string
	Hero_card7    string
	Hero_card8    string
	Hero_card9    string
	Hero_card10   string
}

var ConfHeros []Hero

func GetHeroById(hId string) *Hero {
	for _, h := range ConfHeros {
		if h.Hero_id == hId {
			return &h
		}
	}
	return nil
}
