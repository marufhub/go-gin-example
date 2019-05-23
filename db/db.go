package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	// db.AutoMigrate(&todoModel{})
}

// GetDB Return db
func GetDB() *gorm.DB {
	return db
}
