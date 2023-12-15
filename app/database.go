package app

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	port := Config("DB_PORT")
	host := Config("DB_HOST")
	user := Config("DB_USER")
	password := Config("DB_PASSWORD")
	dbname := Config("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname) // every database has its own dsn
	dialect := mysql.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)

		os.Exit(1)
	}

	return db
}
