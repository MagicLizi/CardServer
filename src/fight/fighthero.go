package fight

import "../config"

type FHero struct {
	Name       string
	CurHp      int
	CurBelief  int
	CurFV      int
	StaticData config.Hero
}
