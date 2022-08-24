package databases

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"user-management/utils"
)

var DB *gorm.DB

func Connect() {
	dbUsername := utils.GoDotEnvVariable("DB_USERNAME")
	dbPass := utils.GoDotEnvVariable("DB_PASSWORD")
	dbHost := utils.GoDotEnvVariable("DB_HOST")
	dbPort := utils.GoDotEnvVariable("DB_PORT")
	dbName := utils.GoDotEnvVariable("DB_NAME")
	dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Connect db fail")
	}

	DB = db
	log.Println("Connect db success")
}
