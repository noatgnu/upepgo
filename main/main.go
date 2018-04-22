package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"upepgo/refseqdb"
	"upepgo/models"
	"time"
	"flag"
	"fmt"
)

var driver = flag.String("driver", "postgres", "Name of database driver to be used")
var dbName = flag.String("db", "", "Name of database")
var user = flag.String("user", "", "Database username")
var pass = flag.String("pass", "", "User password")
var sslmode = flag.String("ssl", "", "SSL Mode")

func main()  {
	flag.Parse()
	setting := fmt.Sprintf("dbname=%v user=%v sslmode=%v", *dbName, *user, *sslmode)
	db, err := sql.Open(*driver, setting)
	if err != nil {
		log.Fatalln(err)
	}
	refseqdb.Truncate(db)
	var refseq models.UpepRefSeqDB
	refseq.Name = "Complete"
	refseq.Insert(db)
	ti := time.Now()
	refseqdb.ReadRefSeqDB("temp/complete.1034.rna.gbff.gz", refseq.ID, db)
	log.Println(time.Since(ti))
	defer db.Close()
}
