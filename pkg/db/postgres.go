package db

import (
	"fmt"

	"github.com/diazharizky/go-graphql-bootstrap/config"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	config.Global.SetDefault("db.postgres.host", "0.0.0.0")
	config.Global.SetDefault("db.postgres.port", "5432")
	config.Global.SetDefault("db.postgres.user", "root")
	config.Global.SetDefault("db.postgres.password", "")
	config.Global.SetDefault("db.postgres.dbname", "")
	config.Global.SetDefault("db.postgres.sslmode", "disable")
}

type postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

func NewPostgres() postgres {
	return postgres{
		Host:     config.Global.GetString("db.postgres.host"),
		Port:     config.Global.GetString("db.postgres.port"),
		User:     config.Global.GetString("db.postgres.user"),
		Password: config.Global.GetString("db.postgres.password"),
		DbName:   config.Global.GetString("db.postgres.dbname"),
		SslMode:  config.Global.GetString("db.postgres.sslmode"),
	}
}

func (pg postgres) Connect() (*gorm.DB, error) {
	cfg := &gorm.Config{}

	db, err := gorm.Open(driver.Open(pg.dsn()), cfg)
	if err != nil {
		return nil, fmt.Errorf("error unable to connect to PostgreSQL DB: %v", err)
	}

	return db, nil
}

func (pg postgres) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg.Host,
		pg.Port,
		pg.User,
		pg.Password,
		pg.DbName,
		pg.SslMode,
	)
}
