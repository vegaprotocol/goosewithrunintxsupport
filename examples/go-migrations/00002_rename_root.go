package main

import (
	"github.com/pressly/goose/v3/internal"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx internal.GooseTx) error {
	_, err := tx.Exec("UPDATE users SET username='admin' WHERE username='root';")
	if err != nil {
		return err
	}
	return nil
}

func Down00002(tx internal.GooseTx) error {
	_, err := tx.Exec("UPDATE users SET username='root' WHERE username='admin';")
	if err != nil {
		return err
	}
	return nil
}
