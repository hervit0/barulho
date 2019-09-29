package main

import (
	"github.com/hervit0/barulho/persistence"
	"github.com/hervit0/barulho/repository"
	"github.com/hervit0/barulho/server"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	db := persistence.Connect()
	persistence.Migrate(db)
	//persistence.Seed(db)

	videosRepo := repository.Video{
		Db: db,
	}
	videosResult := videosRepo.FindBySongName("Ranchera")
	log.Printf("%+v", videosResult)

	server.Do()
}
