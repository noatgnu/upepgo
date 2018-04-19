package upepmodel

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Organism struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null"`
}

type GIAndAccessionLink struct {
	gorm.Model
	GI int `gorm:"int;not null"`
	Accession string `gorm:"type:varchar(20);unique_index:acc"`
}

type Accession struct {
	gorm.Model
	GIAndAccessionLink GIAndAccessionLink `gorm:"foreignkey:GIAcessionID;association_foreignkey:ID"`
	Organism Organism `gorm:"foreignkey:OrganismID;association_foreignkey:ID"`
	GIAccessionID uint
	OrganismID uint
	Position int `gorm:"int;not null"`
	FileP string `gorm:"varchar(1000);not null"`
}

func InitiateDB(db *gorm.DB) {
	db.DropTableIfExists(&Accession{}, &Organism{}, &GIAndAccessionLink{})
	db.AutoMigrate(&Organism{}, &GIAndAccessionLink{}, &Accession{})
	db.Model(&GIAndAccessionLink{}).RemoveIndex("acc")
	o := Organism{Name:"Test"}
	g := GIAndAccessionLink{GI:1,Accession:"214"}
	db.Create(&g)
	db.Create(&o)
	a := Accession{GIAccessionID:g.ID, OrganismID:o.ID, Position:1, FileP:"f:/"}
	db.Create(&a)
	log.Println(a)
	db.Model(&GIAndAccessionLink{}).AddUniqueIndex("acc", "accession")
}