package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/noatgnu/upepgo/refseqdb"
	"flag"
		"github.com/gorilla/mux"
	"net/http"
	"time"
		"github.com/noatgnu/upepgo/helper"
	"os"
	"os/signal"
	"context"
		"github.com/gorilla/handlers"
	"github.com/noatgnu/upepgo/codons"
	"github.com/noatgnu/upepgo/sorf"
	"github.com/noatgnu/upepgo/organisms"
	"path/filepath"
	"bufio"
	"github.com/noatgnu/upepgo/models"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
	"encoding/json"
	"strconv"
	"github.com/jlaffaye/ftp"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strings"
	"github.com/rs/xid"
	"github.com/noatgnu/upepgo/wrapper"
)

var driver = flag.String("driver", "postgres", "Name of database driver to be used")
var dbName = flag.String("db", "", "Name of database")
var user = flag.String("user", "", "Database username")
var pass = flag.String("pass", "", "User password")
var sslmode = flag.String("ssl", "", "SSL Mode")
var port = flag.Int("port", 5432, "Database port")
var runMode = flag.Int("mode", 0, "0 - WebServer \n 1 - Download DBs \n 2 - Process DBs \n 3 - Download and Process DBs \n 4 - Show DB Version \n 5 - InitCodons")
var initCodon = flag.Bool("codon", false, "Populate Starting and Ending Codons Table")
var db *sql.DB
var settings = flag.String("settings", "./settings.json", "Setting File")
var config helper.Settings

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		refseqdb.RequestRefSeqInformationStatus(w, db, false)
		return
	}
}

func AdminRefSeqHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if r.Method == "GET" {
		if vars["origin"] == "remote" {
			refseqdb.RequestRefSeqInformationStatus(w, db, true)
		} else if vars["origin"] == "local"{
			refseqdb.RequestRefSeqInformationStatus(w, db, false)
		}
		return
	}
}

func AdminInitCodons(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		refseqdb.InitCodons(db)
	}
}

func GetSorfHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if r.Method == "GET" {
		sorf.GetAllSorf(w, vars)
	}
}

func GetCodonsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		codons.GetAllCodons(w)
	}
}

func GetOrganismsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if r.Method == "GET" {
		organisms.GetAllOrganisms(w, vars)
	}
}

func GetRefSeqDB(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		refseqdb.GetLocalDB(w, db)
		return
	}
}

func AdminRefreshRefSeqHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		vars := mux.Vars(r)
		totalNumber, err := strconv.Atoi(vars["totalNumber"])
		if err != nil {
			log.Panicln(err)
		}
		client := refseqdb.CreateFtpClient()
		var fileList []ftp.Entry
		json.NewDecoder(r.Body).Decode(&fileList)
		if totalNumber > 0 {
			ti := time.Now()
			test, err := strconv.ParseBool(vars["test"])
			if err != nil {
				log.Panicln(err)
			}
			startingCodons, endingCodons := refseqdb.GetAllCodons(db)
			eMap := make(map[string]*models.UpepCodon)
			for _, v := range endingCodons {
				eMap[v.Sequence] = v
			}
			sMap := make(map[string]*models.UpepCodon)
			for _, v := range startingCodons {
				sMap[v.Sequence] = v
			}
			version, err := strconv.Atoi(vars["version"])
			rsdb := models.UpepRefSeqDB{Name: vars["dbname"], Version: version}
			if err != nil {
				log.Panicln(err)
			}
			exists, err := models.UpepRefSeqDBS(db, qm.Where("name=?", vars["dbname"])).Exists()
			if err != nil {
				log.Panicln(err)
			}
			if exists {
				log.Printf("Deleting DB %s version %v", vars["dbname"], version)
				refseqdb.DeleteDB(vars["dbname"], db)
				log.Printf("Finished Deleting DB %s version %v", vars["dbname"], version)
			}
			log.Printf("Creating DB %s version %v", vars["dbname"], version)
			err = rsdb.Insert(db)
			fileMap := make(map[string]*helper.BlastDBWriter)

			writerChan := make(chan []string)
			go func() {
				if totalNumber != 0 {
					for _, val := range fileList {
						totalNumber -= 1
						if strings.HasSuffix(val.Name, "rna.gbff.gz") {

							if _, err := os.Stat(filepath.Join(config.Temp, val.Name)); os.IsNotExist(err) {
								log.Printf("Started downloading %v", val.Name)
								refseqdb.DownloadToTemp(client, val.Name, refseqdb.ReleaseMap[vars["dbname"]]+"/")
								log.Printf("Finished downloading %v", val.Name)
							}
							log.Printf("Started processing %v", val.Name)
							refseqdb.ReadRefSeqDB(filepath.Join(config.Temp, val.Name), rsdb.ID, db, sMap, eMap, writerChan)
							log.Printf("Finished processing %v", val.Name)
							if test {
								break
								close(writerChan)
							}
						}
					}
				} else {
					close(writerChan)
				}
			}()
			for e := range writerChan {
				fileName := fmt.Sprintf("%v_%v_%v", refseqdb.ReleaseMap[vars["dbname"]], e[1], e[2])

				if _, ok := fileMap[fileName]; !ok {
					filePath := filepath.Join(config.Temp, fileName)
					f, err := os.Create(filePath)
					if err != nil {
						log.Panicln(err)
					}
					w := bufio.NewWriter(f)
					start, err := strconv.ParseInt(e[1], 10, 64)
					if err != nil {
						log.Panicln(err)
					}
					end, err := strconv.ParseInt(e[2], 10, 64)
					if err != nil {
						log.Panicln(err)
					}


					bdb := models.UpepBlastDB{
						Title:           vars["dbname"],
						Path:            filepath.Join(config.Blastdb, fileName),
						Description:     "",
						UpepRefSeqDBID:  rsdb.ID,
						StartingCodonID: start,
						EndingCodonID:   end,
					}
					exists, err := models.UpepBlastDBS(db, qm.Where("path=?", bdb.Path)).Exists()
					if err != nil {
						log.Panicln(err)
					}
					if exists {
						blastdbs, err := models.UpepBlastDBS(db, qm.Where("path=?", bdb.Path)).All()
						if err != nil {
							log.Panicln(err)
						}
						blastdbs.DeleteAll(db)
					}

					err = bdb.Insert(db)
					if err != nil {
						log.Panicln(err)
					}
					b := helper.BlastDBWriter{File: f, Writer: w, BlastDB: &bdb}
					fileMap[fileName] = &b
				}
				fileMap[fileName].Writer.WriteString(e[0])
			}
			for k, v := range fileMap {
				v.Writer.Flush()
				v.File.Close()
				refseqdb.MakeBlastDB(config.MakeBlastDB, filepath.Join(config.Temp, k), v.BlastDB.Path)
			}

			log.Printf("Finished Creating DB %s version %v", vars["dbname"], version)
			log.Printf("Completed in %v", time.Since(ti))
		}
		return
	}
}

func GetBlastDBHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if r.Method == "GET" {
		id, err := strconv.ParseInt(vars["blastdbId"], 10, 64)
		if err != nil {
			log.Panicln(err)
		}
		if id == 0 {
			blastdbs, err := models.UpepBlastDBS(db, qm.Select("id", "upep_ref_seq_db_id", "starting_codon_id", "ending_codon_id")).All()
			if err != nil {
				log.Panicln(err)
			}
			encoder := json.NewEncoder(w)
			err = encoder.Encode(blastdbs)
			if err != nil {
				log.Panicln(err)
			}
			return
		}
	}
}

func SearchSORFHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var query helper.SearchQuery
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&query)
		if err != nil {
			log.Panicln(err)
		}
		evalue, err := strconv.ParseFloat(query.Evalue, 64)
		if err != nil {
			log.Panicln(err)
		}
		guid := xid.New()
		inp := filepath.Join(config.Temp, guid.String()+".in")
		f, err := os.Create(inp)
		writer := bufio.NewWriter(f)
		queryMap := make(map[string]*helper.Sequence)
		for s := range query.Sequences {
			queryMap[query.Sequences[s].Header] = &query.Sequences[s]
			writer.WriteString(query.Sequences[s].ToString())
		}
		writer.Flush()
		f.Close()
		for b := range query.BlastDB {
			blastDB, err := models.FindUpepBlastDB(db, query.BlastDB[b])
			if err != nil {
				log.Panicln(err)
			}
			tblastcmd := wrapper.TBlastXCommandline{
				Command:     config.Blast,
				In:          inp,
				DB:          blastDB.Path,
				Filter:      false,
				Out:         inp+".xml",
				Evalue:      evalue,
				Outfmt:      5,
				CommandLine: nil,
			}
			tblastcmd.Execute()
			queriesOut := wrapper.ParseHitTBlastXOutputXML(inp+".xml", queryMap, db)
			log.Println(queriesOut)
		}
		log.Println(evalue)
		log.Println(query.Sequences)
		log.Println(query.BlastDB)
	}
}

func main()  {
	var err error
	flag.Parse()
	if *settings != "" {
		f, err := os.Open(*settings)
		if err != nil {
			log.Panicln(err)
		}
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&config)
	} else {
		config.DBDriver = *driver
		config.DBName = *dbName
		config.DBUser = *user
		config.DBPort = *port
		config.DBSSL = *sslmode
		config.DBPass = *pass
		config.DBRunmode = *runMode
	}

	setting := fmt.Sprintf("dbname=%v user=%v sslmode=%v password=%v port=%v host=%v", config.DBName, config.DBUser, config.DBSSL, config.DBPass, config.DBPort, config.DBServer)
	db, err = sql.Open(config.DBDriver, setting)
	if err != nil {
		log.Fatalln(err)
	}
	boil.SetDB(db)
	defer db.Close()
	switch config.DBRunmode {
	case 0:
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-type", "Access-Control-Allow-Origin"})
		originsOk := handlers.AllowedOrigins([]string{"http://localhost:4300", "http://localhost:4200"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		r := mux.NewRouter()
		r.HandleFunc("/admin/", AdminHandler)
		r.HandleFunc("/admin/request/refseq/{origin}", AdminRefSeqHandler)
		r.HandleFunc("/request/sorf/limit/{limit}/offset/{offset}/startingCodon/{startingCodon}/stoppingCodon/{stoppingCodon}/organism/{organism}/refseqname/{refseqname}/dbId/{dbId}/", GetSorfHandler)
		r.HandleFunc("/request/codons/", GetCodonsHandler)
		r.HandleFunc("/request/refseqdb/", GetRefSeqDB)
		r.HandleFunc("/request/organisms/limit/{limit}/{dbId}/", GetOrganismsHandler)
		r.HandleFunc("/admin/request/refreshdb/{dbname}/{totalNumber}/{version}/{test}/", AdminRefreshRefSeqHandler)
		r.HandleFunc("/admin/request/initcodons/", AdminInitCodons)
		r.HandleFunc("/request/blastdb/{blastdbId}/", GetBlastDBHandler)
		r.HandleFunc("/request/search-sorf/", SearchSORFHandler)
		srv := &http.Server{
			Addr:         "0.0.0.0:8080",
			// WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler: handlers.CORS(headersOk, originsOk, methodsOk)(r),
		}
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}()
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
		defer cancel()
		srv.Shutdown(ctx)
		log.Println("shutting down")
		os.Exit(0)


	case 5:
		helper.DownMigrations(db)
		helper.UpMigrations(db)
		refseqdb.InitCodons(db)
	}


}
