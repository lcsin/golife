package database

import (
	"log"

	"github.com/golife/gin-layout/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	if config.App().Database.Driver == "mysql" {
		DB, err = gorm.Open(mysql.Open(config.App().Database.Source), &gorm.Config{})
		if err != nil {
			log.Fatalf("cannot connect mysql: %v", err)
		}
		log.Println("connect mysql success ...")
	}
}

func Close() error {
	d, err := DB.DB()
	if err != nil {
		return err
	}
	return d.Close()
}
