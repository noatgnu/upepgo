package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"log"
	"upepgo/upepmodel"
)


func main()  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	upepmodel.InitiateDB(db)

}
