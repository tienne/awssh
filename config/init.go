package config

import (
	"fmt"
	"github.com/tienne/awssh/database"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

var (
	EnvPrefix    = "__" + AppName
	ConfigHome   = path.Join(os.Getenv("HOME"), "."+AppName)
	DBPath       = filepath.Join(ConfigHome, database.FileName)
	Dir          = filepath.Join(ConfigHome, "aws")
	KeysDir      = filepath.Join(ConfigHome, "keys")
	FirstInstall bool
)

func init() {
	os.Setenv(EnvPrefix+"_HOME", ConfigHome)
	os.Setenv(EnvPrefix+"_CACHE", filepath.Join(ConfigHome, "cache"))
	os.Setenv(EnvPrefix+"_KEYS_DIR", KeysDir)
}

func InitAwlessEnv() error {
	_, err := os.Stat(DBPath)

	FirstInstall = os.IsNotExist(err)
	os.Setenv(EnvPrefix+"_FIRST_INSTALL", strconv.FormatBool(FirstInstall))

	os.MkdirAll(KeysDir, 0700)

	if FirstInstall {
		fmt.Fprintln(os.Stderr, AsciiLogo)
		fmt.Fprintln(os.Stderr, "Welcome! Resolving environment data...")
		fmt.Fprintln(os.Stderr)

		if err = InitConfig(resolveRequiredConfigFromEnv()); err != nil {
			return err
		}

		err = database.Execute(func(db *database.DB) error {
			return db.SetStringValue("current.version", Version)
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot store current version in db: %s\n", err)
		}
	}

	if err = LoadConfig(); err != nil {
		return err
	}

	return nil
}
