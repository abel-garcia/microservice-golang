package model

type Book struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func (db *DB) SaveBooks() ([]*Book, error) {
	bks := make([]*Book, 0)
	return bks, nil
}
