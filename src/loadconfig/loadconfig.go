package loadconfig

import (
	"encoding/json"
	"golang-exec/model"
	"io/ioutil"
)

func LoadConfig() []model.CommandList {
	var CommandList []model.CommandList
	readFile, _ := ioutil.ReadFile("../commands.json")
	json.Unmarshal(readFile, &CommandList)
	return CommandList
}
