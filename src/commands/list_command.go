package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/seanervinson/vc/models"
	"github.com/seanervinson/vc/utils"
)

type ListCommand struct {
	Verbose bool
}

const CodeHeader = "CODE"
const DescriptionHeader = "DESCRIPTION"
const PathHeader = "PATH"

func (command ListCommand) Execute() {
	data, err := utils.LoadFile(configPath)
	if err != nil {
		os.Exit(1)
	}
	var configs []models.Config
	if err := json.Unmarshal(data, &configs); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := "%-16v%-40v\n"
	if command.Verbose {
		s = "%-16v%-24v%-40v\n"
		fmt.Printf(s, CodeHeader, DescriptionHeader, PathHeader)
	} else {
		fmt.Printf(s, CodeHeader, DescriptionHeader)
	}
	for _, config := range configs {
		if command.Verbose {
			description := ""
			if config.Description != nil {
				description = *config.Description
			}
			fmt.Printf(s, config.Code, description, config.Path)
		} else {
			fmt.Printf(s, config.Code, config.Path)
		}
	}
}
