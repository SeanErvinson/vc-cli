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
	data, err := utils.LoadFile(configPath)
	var configs []models.Config
	if err == nil {
		json.Unmarshal(data, &configs)
	}

	if codeExists(configs, command) {
		fmt.Printf("Code '%v' already exists.\n", command.Code)
		os.Exit(0)
	}

	newConfig := models.Config{Description: command.Description, Code: command.Code}
	directoryPath := getProjectPath(command.ProjectPath)
	if command.Workspace {
		filepath.Walk(directoryPath, func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == ".code-workspace" {
				directoryPath = path
			}
			return nil
		})
	}
	newConfig.Path = directoryPath
	configs = append(configs, newConfig)
	data, _ = json.Marshal(configs)
	utils.SaveFile(configPath, data)
}

func getProjectPath(path string) string {
	cd, _ := os.Getwd()
	var directoryPath = cd
	if path != "." {
		exist, err := utils.IsDirectoryExist(path)
		if err != nil || !exist {
			fmt.Println("Directory does not exist.")
			os.Exit(0)
		}
		directoryPath = cd
	}
	return directoryPath
}

func codeExists(configs []models.Config, command SetCommand) bool {
	for _, config := range configs {
		if config.Code == command.Code {
			return true
		}
	}
	return false
}
