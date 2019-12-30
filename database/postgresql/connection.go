package postgresql

import (
	_ "github.com/lib/pq"
	"database/sql"
	"framework/app/model"
	"log"
)

func Conn() *model.DB {
	if db, err := Connection("postgres://user:pass@localhost/bookstore"); err == nil {
		return db
	} else {
		log.Panic(err)
	}
	return nil
}

func Connection(dataSourceName string) (*model.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &model.DB{db}, nil
}
