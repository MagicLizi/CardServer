package config

type Card struct {
	Card_id       string
	Card_name     string
	Card_skin     string
	Card_describe string
	Card_type     string
	Card_hp       int
	Card_cost     int
	Card_combat   int
	Card_tenet    int
	Card_skill1   string
	Card_skill2   string
}

var ConfCards []Card
