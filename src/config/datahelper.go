package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Card struct {
	Card_id string
	Card_name string
	Card_skin string
	Card_describe string
	Card_type string
	"card_hp": "卡牌生命",
	"card_cost": "卡牌信仰费用",
	"card_combat": "提供战力",
	"card_tenet": "提供信仰",
	"card_skill1": "卡牌技能1",
	"card_skill2": "卡牌技能2"
}

var CardsData []Card

func InitDataConfig() {
	loadJSONFile("card", &CardsData)
}

func loadJSONFile(filename string, v interface{}) {
	log.Println("load json file " + filename)
	data, err := ioutil.ReadFile("D:/WorkCode/CardServer/data/" + filename)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Println(err)
		return
	}
}
