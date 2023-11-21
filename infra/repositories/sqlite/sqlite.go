package sqlite

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func BuildDataBase(ctx context.Context, db *sql.DB) {
	log.Println("seed database")

	_, err := db.ExecContext(
		ctx,
		`
		CREATE TABLE IF NOT EXISTS operations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			symbol VARCHAR(255) NOT NULL,
			type VARCHAR(255) NOT NULL,
			price INTEGER NOT NULL
		);
		`,
	)

	if err != nil {
		log.Fatalf("cant seed operation database tabele: %v", err)
	}

	log.Println("Finish seed")
}
