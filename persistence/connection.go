package persistence

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"time"
)

func Connect() *gorm.DB {
	host := os.Getenv("DATABASE_HOST") // localhost if outside docker
	dbConnection := fmt.Sprintf("host=%v user=example password=example DB.name=example port=5432 sslmode=disable", host)

	db, err := gorm.Open("postgres", dbConnection)
	if err != nil {
		log.Print(err)
		return nil
	}
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(7)
	//db.Set("gorm:auto_preload", true)

	return db
}
