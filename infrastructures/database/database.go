package database

import (
	"SmartAerators/infrastructures/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	DatabaseConf() *gorm.DB
}

type database struct {
}

func NewDatabase() *database {
	return &database{}
}

func (d *database) DatabaseConf() *gorm.DB {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := "host=" + config.Database.Host + " user=" + config.Database.Username + " password=" + config.Database.Password + " dbname=" + config.Database.Name + " port=" + config.Database.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	// digunakan untuk auto migrate dari entity ke database
	// db.AutoMigrate(&models.User{})

	return Db
}
