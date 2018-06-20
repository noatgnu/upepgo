package helper

import (
	"github.com/rubenv/sql-migrate"
	"database/sql"
	"log"
	"fmt"
		"github.com/jlaffaye/ftp"
	"github.com/noatgnu/upepgo/models"
	"os"
	"bufio"
)

type SearchQuery struct {
	BlastDB []int64 `json:"Blastdb"`
	Evalue string `json:"Evalue"`
	Sequences []Sequence `json:"Sequences"`
}

type Sequence struct {
	Header string `json:"Header"`
	Seq string `json:"Seq"`
}

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
	TBlastX string
	DBDriver string
	DBName string
	DBUser string
	DBPass string
	DBSSL string
	DBPort int
	DBRunmode int
	DBServer string
	LaganLoc string
}

type BlastDBWriter struct {
	File *os.File
	Writer *bufio.Writer
	BlastDB *models.UpepBlastDB
}

type LaganAlignment struct {
	Target string
	MidLine string
	Query string
}

var migrations = &migrate.FileMigrationSource{
	Dir: "migrations/postgres",
}

func (s Sequence) ToString() string {
	return fmt.Sprintf(">%v\n%v\n\n", s.Header, s.Seq)
}

func UpMigrations(db *sql.DB) {
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	fmt.Println("Create Schema and Table.")
}

func DownMigrations(db *sql.DB) {
	n, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
	fmt.Println("Delete Schema and Table.")
}