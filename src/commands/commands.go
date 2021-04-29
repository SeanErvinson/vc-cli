package commands

import (
	"os"
	"path/filepath"
)

type Action interface {
	Execute()
}

const configFile = ".vc.config.json"

var homeDir, _ = os.UserHomeDir()
var configPath = filepath.Join(homeDir, configFile)
