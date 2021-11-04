package db

import (
	"log"
	"github.com/httpmon/user/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(config config.Database) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Cstring()), &gorm.Config{})
	if err != nil {
		log.Printf("can not open connection to database due to the following err\n: %s", err)
	}

	return db
}
