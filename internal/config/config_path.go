package config

import (
	"os"
	"path/filepath"
	"runtime"
)

const (
	RF_CONFIG_DIR   = "GH_CONFIG_DIR"
	XDG_CONFIG_HOME = "XDG_CONFIG_HOME"
	APP_DATA        = "AppData"
)

var Dir = getDir()
var Path = filepath.Join(Dir, "config.toml")

// getDir return folder that contains config
func getDir() string {
	var path string
	if a := os.Getenv(RF_CONFIG_DIR); a != "" {
		path = a
	} else if b := os.Getenv(XDG_CONFIG_HOME); b != "" {
		path = filepath.Join(b, "rf")
	} else if c := os.Getenv(APP_DATA); runtime.GOOS == "windows" && c != "" {
		path = filepath.Join(c, "RF CLI")
	} else {
		d, _ := os.UserHomeDir()
		path = filepath.Join(d, ".config", "rf")
	}

	return path
}

func dirExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.IsDir()
}

func fileExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir()
}
