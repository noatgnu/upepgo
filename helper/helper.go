package helper

import (
	"github.com/rubenv/sql-migrate"
	"database/sql"
	"log"
	"fmt"
)

var migrations = &migrate.FileMigrationSource{
	Dir: "migrations",
}

func UpMigrations(db *sql.DB) {
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Create Schema and Table.")
}

func DownMigrations(db *sql.DB) {
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Delete Schema and Table.")
}