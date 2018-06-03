package helper

import (
	"github.com/rubenv/sql-migrate"
	"database/sql"
	"log"
	"fmt"
		"github.com/jlaffaye/ftp"
	"upepgo/models"
	"os"
	"bufio"
)

type RefSeqQuery struct {
	CurrentDB []*models.UpepRefSeqDB
	FileListMap map[string][]*ftp.Entry
	RemoteDB int
	ReleaseMap map[string]string
}

type Settings struct {
	Temp string
	Blastdb string
	Blast string
	MakeBlastDB string
	DBDriver string
	DBName string
	DBUser string
	DBPass string
	DBSSL string
	DBPort int
	DBRunmode int
}

type BlastDBWriter struct {
	File *os.File
	Writer *bufio.Writer
	BlastDB *models.UpepBlastDB
}

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