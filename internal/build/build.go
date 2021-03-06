package build

import "runtime/debug"

var (
	Version = "DEV"
	Date    = "NOT_SET"
)

func init() {
	if Version != "DEV" {
		return
	}

	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
		Version = info.Main.Version
	}
}
