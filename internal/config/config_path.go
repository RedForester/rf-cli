package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	RfConfigDir   = "GH_CONFIG_DIR"
	XdgConfigHome = "XDG_CONFIG_HOME"
	AppData       = "AppData"

	Dir      = ".config"
	FileName = "rf"
	FileExt  = "yml"
)

var CurrentPath = ""

func GetConfigFile() string {
	if CurrentPath != "" {
		return CurrentPath
	}
	return fmt.Sprintf("%s/%s.%s", GetConfigHome(), FileName, FileExt)
}

func GetConfigHome() string {
	var path string
	if a := os.Getenv(RfConfigDir); a != "" {
		path = a
	} else if b := os.Getenv(XdgConfigHome); b != "" {
		path = b
	} else if c := os.Getenv(AppData); runtime.GOOS == "windows" && c != "" {
		path = c
	} else {
		d, _ := os.UserHomeDir()
		path = filepath.Join(d, Dir)
	}

	return path
}
