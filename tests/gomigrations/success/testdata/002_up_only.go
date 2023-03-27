package gomigrations

import (
	"github.com/pressly/goose/v3/internal"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(up002, nil)
}

func up002(tx internal.GooseTx) error {
	q := "INSERT INTO foo VALUES (1, 1, 'Alice')"
	_, err := tx.Exec(q)
	return err
}
