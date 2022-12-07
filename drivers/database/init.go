package database

import (
	"github.com/asdine/storm/v3"
	"log"
)

func InitDatabse() *storm.DB {
	db, err := storm.Open("database/my.db")
	if err != nil {
		log.Fatalln("Database open error.")
	}

	if err == nil {
		log.Println("Database is ok !")
	}

	defer db.Close()

	return db
}
