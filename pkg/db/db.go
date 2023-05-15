package db

import (
	"fmt"

	"github.com/diazharizky/go-graphql-bootstrap/config"
	"github.com/diazharizky/go-graphql-bootstrap/internal/enum"
	"gorm.io/gorm"
)

var (
	dbType enum.DbType
)

func init() {
	config.Global.SetDefault("db.type", "postgres")

	dbType = enum.DbType(config.Global.GetString("db.type"))
}

func GetConnection() (conn *gorm.DB, err error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	conn, err = db.Connect()
	if err != nil {
		return nil, err
	}

	return
}

func getDB() (db IDatabase, err error) {
	if !dbType.Validate() {
		err = fmt.Errorf("error invalid database type: %s", dbType.String())
		return
	}

	switch dbType {
	case enum.DbTypePostgres:
		db = NewPostgres()
	default:
		db = NewPostgres()
	}

	return
}
