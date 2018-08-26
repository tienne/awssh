package database

import (
	"fmt"
	"github.com/golang/leveldb"
	"github.com/golang/leveldb/db"
	"github.com/tienne/awssh/config"
)

var (
	FileName = config.AppName + ".db"
)

type DB struct {
	db *leveldb.DB
}

func open(path string) (*DB, error) {
	db, err := leveldb.Open(path, &db.Options{})
	if err != nil {
		return nil, fmt.Errorf("opening db at %s: %s (any awssh existing process running?)", path, err)
	}
	return &DB{db: db}, nil
}
