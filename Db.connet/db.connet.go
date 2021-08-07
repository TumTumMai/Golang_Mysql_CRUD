package Dbconnet

import (
	"goec/model"
	"log"

	"github.com/jinzhu/gorm"
)

type DB struct {
	DB *gorm.DB
}

var database *gorm.DB

//ให้เชื่อมต่อฐานข้อมูลเมื่อ Initialize
func Initialize() {
	db, err := gorm.Open("mysql", "root:helloworld@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Item{})

	database = db
}

func GetDatabase() *gorm.DB {
	return database
}
