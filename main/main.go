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
var port = flag.Int("port", 5432, "Database port")
var runMode = flag.Int("mode", 0, "0 - WebServer \n 1 - Download DBs \n 2 - Process DBs \n 3 - Download and Process DBs \n 4 - Show DB Version")
var initCodon = flag.Bool("codon", false, "Populate Starting and Ending Codons Table")

func main()  {
	flag.Parse()
	setting := fmt.Sprintf("dbname=%v user=%v sslmode=%v password=%v port=%v", *dbName, *user, *sslmode, *pass, *port)
	db, err := sql.Open(*driver, setting)
	if err != nil {
		log.Fatalln(err)
	}
	refseqdb.Truncate(db)
	var refseq models.UpepRefSeqDB
	if *initCodon {
		refseqdb.InitCodons(db)
	}
	refseq.Name = "Complete"
	refseq.Insert(db)
	ti := time.Now()
	startingCodons, endingCodons := refseqdb.GetAllCodons(db)
	eMap := make(map[string]*models.UpepCodon)
	for _,v := range endingCodons {
		eMap[v.Sequence] = v
	}
	refseqdb.ReadRefSeqDB("temp/complete.1034.rna.gbff.gz", refseq.ID, db, startingCodons, eMap)
	log.Println(time.Since(ti))
	defer db.Close()
}
