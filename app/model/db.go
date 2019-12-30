package model

import "database/sql"

type Datastore interface {
	SaveBooks() ([]*Book, error)
}

type DB struct {
	*sql.DB
}
