package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func InitDataConfig() {
	loadJSONFile("card", &ConfCards)
	loadJSONFile("centercard", &ConfCenterCards)
	loadJSONFile("player1card", &P1Cards)
	loadJSONFile("player2card", &P2Cards)
	loadJSONFile("hero", &ConfHeros)
	loadJSONFile("skill", &ConfSkills)
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
