package commands

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/seanervinson/vc/models"
	"github.com/seanervinson/vc/utils"
)

type CodeCommand struct {
	Code string
}

func (command CodeCommand) Execute() {
	data, _ := utils.LoadFile(configPath)
	var configs []models.Config
	if err := json.Unmarshal(data, &configs); err != nil {
		fmt.Println(err)
	}
	for _, config := range configs {
		if path := config.Path; config.Code == command.Code {
			cmd := exec.Command("code", path)
			_, err := cmd.Output()
			if err != nil {
				fmt.Println("Something unexpected happened.")
			}
			break
		}
	}
}
