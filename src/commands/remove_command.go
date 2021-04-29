package commands

import (
	"encoding/json"

	"github.com/seanervinson/vc/models"
	"github.com/seanervinson/vc/utils"
)

type RemoveCommand struct {
	Code string
}

func (command RemoveCommand) Execute() {
	data, _ := utils.LoadFile(configPath)
	var configs []models.Config
	json.Unmarshal(data, &configs)
	for i, config := range configs {
		if config.Code == command.Code {
			configs = append(configs[:i], configs[i+1:]...)
		}
	}
	data, _ = json.Marshal(configs)
	utils.SaveFile(configPath, data)
}
