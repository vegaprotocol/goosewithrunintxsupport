package gomigrations

import (
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(up001, down001)
}

func up001(tx goose.Tx) error {
	q := "CREATE TABLE foo (id INT, subid INT, name TEXT)"
	_, err := tx.Exec(q)
	return err
}

func down001(tx goose.Tx) error {
	q := "DROP TABLE IF EXISTS foo"
	_, err := tx.Exec(q)
	return err
}
