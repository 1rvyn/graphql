package database

import (
	"errors"
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

	err = db.AutoMigrate(&models.Employee{}, &models.Department{})
	if err != nil {
		return
	}

	Database = Dbinstance{Db: db}

	// Create some initial departments
	departments := []models.Department{
		{Name: "Software", Description: "This is the software department"},
		{Name: "Human Resources", Description: "This is the human resources department"},
		{Name: "Personel", Description: "This is the personel department"},
	}

	for _, dept := range departments {
		// Check if the department already exists
		var existingDept models.Department
		result := db.First(&existingDept, "name = ?", dept.Name)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// If the department doesn't exist, create it
			if err := db.Create(&dept).Error; err != nil {
				log.Println("Failed to create department:", err)
			} else {
				log.Println("Created department:", dept.Name)
			}
		}
	}
}
