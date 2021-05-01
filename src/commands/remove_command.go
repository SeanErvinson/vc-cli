package commands

import (
	"encoding/json"
	"os"

	"github.com/seanervinson/vc/models"
	"github.com/seanervinson/vc/utils"
)

type RemoveCommand struct {
	Code string
}

func (command RemoveCommand) Execute() {
	data, err := utils.LoadFile(configPath)
	if err != nil {
		os.Exit(1)
	}
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
