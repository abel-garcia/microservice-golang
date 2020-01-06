package model

import (
	pg "github.com/jackc/pgx"
)


type DBpg struct {
	*pg.Conn
}

func (d *DBpg) GetConnection() *pg.Conn {
	return d.Conn
}