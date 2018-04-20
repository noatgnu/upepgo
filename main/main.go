package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"log"
	"upepgo/upepmodel"
)
var DB *gorm.DB
func init() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database successfully connected")
}


func main()  {
	defer DB.Close()
	upepmodel.InitiateDB(DB)

}
