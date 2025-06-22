package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JKasus/go_final_project/pkg/constants"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	if _, err = db.Exec(constants.Schema); err != nil {
		return fmt.Errorf("schema creation error: %w", err)
	}

	return nil
}
