package config

type CenterCard struct {
	Card_id     int
	Card_number int
	Inevitable  int
}

var AllCenterCards []CenterCard

var MustCenterCards []CenterCard

var RandomCenterCards []CenterCard

func ParseCenterCards() {
	for _, v := range AllCenterCards {
		if v.Inevitable == 1 {
			MustCenterCards = append(MustCenterCards, v)
		} else {
			RandomCenterCards = append(RandomCenterCards, v)
		}
	}
}
