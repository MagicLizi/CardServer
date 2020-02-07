package config

type Hero struct {
	Hero_id       string
	Hero_name     string
	Hero_skin     string
	Hero_describe string
	Hero_hp       int
	Hero_card     string
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

func ParseHeros() {

}
