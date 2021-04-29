package commands

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/seanervinson/vc/models"
	"github.com/seanervinson/vc/utils"
)

type SetCommand struct {
	Code        string
	Workspace   bool
	Description *string
	ProjectPath string
}

func (command SetCommand) Execute() {
	newConfig := models.Config{Description: command.Description, Code: command.Code}
	cd, _ := os.Getwd()
	var directoryPath = cd
	if command.ProjectPath != "." {
		exist, err := utils.IsDirectoryExist(command.ProjectPath)
		if err != nil || !exist {
			fmt.Println("Directory does not exist.")
			os.Exit(0)
		}
		directoryPath = cd
	}
	if command.Workspace {
		filepath.Walk(directoryPath, func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == ".code-workspace" {
				directoryPath = path
			}
			return nil
		})
	}
	newConfig.Path = directoryPath
	data, _ := json.Marshal(newConfig)
	utils.SaveFile(configPath, data)
}
