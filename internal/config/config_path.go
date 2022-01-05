package config

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

const (
	RfConfigDir   = "GH_CONFIG_DIR"
	XdgConfigHome = "XDG_CONFIG_HOME"
	AppData       = "AppData"

	Dir         = ".config"
	FileName    = "rf"
	FileNameWin = "RF CLI"
	FileExt     = "yml"
)

func GetConfigHome() (string, error) {
	var path string
	if a := os.Getenv(RfConfigDir); a != "" {
		path = a
	} else if b := os.Getenv(XdgConfigHome); b != "" {
		path = filepath.Join(b, FileName)
	} else if c := os.Getenv(AppData); runtime.GOOS == "windows" && c != "" {
		path = filepath.Join(c, FileNameWin)
	} else {
		d, _ := os.UserHomeDir()
		path = filepath.Join(d, Dir, FileName)
	}

	if FileExists(path) != true {
		return "", errors.New("config not exist")
	}

	return path, nil
}

func FileExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir()
}
