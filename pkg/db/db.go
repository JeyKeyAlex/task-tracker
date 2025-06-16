package db

import (
	"database/sql"
	"fmt"
	"github.com/JKasus/go_final_project/pkg/constants"
	"os"

	_ "modernc.org/sqlite"
)

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}
	defer db.Close()

	if install {
		if _, err = db.Exec(constants.Schema); err != nil {
			return fmt.Errorf("schema creation error: %w", err)
		}
	}

	return nil
}
