package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/1rvyn/graphql-service/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var server = os.Getenv("DB_SERVER") // localhost
var port = 1433
var portStr = strconv.Itoa(port)

var user = os.Getenv("DB_USER")         // sa
var password = os.Getenv("DB_PASSWORD") // your_password
var database = os.Getenv("DB_NAME")     // takehome

var Database Dbinstance

func ConnectDb() {
	sqlserverconn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable;", server, user, password, portStr, database)

	db, err := gorm.Open(sqlserver.Open(sqlserverconn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database \n", err.Error())
		os.Exit(2)
	}

	log.Printf("There was a successful connection to the: %s Database", database)

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// TODO: add migrations

	err = db.AutoMigrate(&models.Employee{}, &models.Department{})
	if err != nil {
		return
	}

	Database = Dbinstance{Db: db}
}
