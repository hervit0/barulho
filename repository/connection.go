package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

func Connect() *gorm.DB {
	dbConnection := "host=localhost user=example password=example DB.name=example port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dbConnection)
	if err != nil {
		log.Print(err)
		return nil
	}
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(7)

	return db
}
