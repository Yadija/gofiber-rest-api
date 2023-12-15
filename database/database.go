package database

import (
	"fmt"
	"gofiber-restapi/config"
	"gofiber-restapi/internals/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// database instance
type DbInstance struct {
	Db *gorm.DB
}

// declare global variable
var DB DbInstance

func OpenConnection() {
	port := config.Config("DB_PORT")
	host := config.Config("DB_HOST")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	dbname := config.Config("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname) // every database has its own dsn
	dialect := mysql.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)

		os.Exit(1)
	}

	log.Println("Connected to database!")
	log.Println("Running migrations...")
	db.AutoMigrate(&model.User{}, &model.Thread{})
	log.Println("Database migrations complete!")

	DB = DbInstance{
		Db: db,
	}
}
