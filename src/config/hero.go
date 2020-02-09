package config

type Hero struct {
	Hero_id       int
	Hero_name     string
	Hero_skin     int
	Hero_describe string
	Hero_hp       int
	Hero_card     string
}

var ConfHeros []Hero

func GetHeroById(hId int) *Hero {
	for _, h := range ConfHeros {
		if h.Hero_id == hId {
			return &h
		}
	}
	return nil
}
