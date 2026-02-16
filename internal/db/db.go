// Package db provides a database interface for the gokedex application.
package db

import (
	"log"

	"github.com/tidwall/buntdb"
)

func NewDB() *buntdb.DB {
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}
