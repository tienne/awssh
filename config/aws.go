package config

import (
	"os"
	"path/filepath"
	"runtime"
)

var AWSRootDir = func() string {
	var home string
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	} else {
		home = os.Getenv("HOME")
	}
	return filepath.Join(home, ".aws")
}
