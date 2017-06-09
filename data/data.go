package data

import (
	"fmt"

	"go-ws/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DataProvider interface {
	Hello() (string, error)
}

type DB struct {
	*sqlx.DB
}

func NewDataProvider(config *config.Config) (*DB, error) {
	connStr := buildConnStr(config)
	db, err := sqlx.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(config.DbMaxOpenConss)
	db.SetMaxIdleConns(config.DbMaxIdleConss)
	return &DB{db}, nil
}

func buildConnStr(config *config.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName)
}
