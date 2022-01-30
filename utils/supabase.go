package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Db *gorm.DB

func ConnectSupabase() {
	var err error
	Db, err = gorm.Open(postgres.Open(os.Getenv("SUPABASEURL")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
}
