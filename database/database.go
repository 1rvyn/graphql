package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/1rvyn/graphql-service/models"
	"github.com/1rvyn/graphql-service/utils"
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

	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ { // 5 retries
		db, err = gorm.Open(sqlserver.Open(sqlserverconn), &gorm.Config{})
		if err == nil {
			break // break the loop if connection is successful
		}
		log.Println("Unable to connect to the database, retrying in 5 seconds...")
		time.Sleep(8 * time.Second) // wait for 5 seconds before retrying
	}

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
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

	// Create an initial employee
	employee := models.Employee{
		FirstName:    "George",
		LastName:     "Test",
		Username:     "george",
		Password:     utils.HashPassword("password"),
		Email:        "george@test.com",
		DOB:          time.Now(), // set a proper date
		DepartmentID: 1,          // ID of the department this employee belongs to
		Position:     "Developer",
	}

	// Check if the employee already exists
	var existingEmployee models.Employee
	result := db.First(&existingEmployee, "username = ?", employee.Username)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// If the employee doesn't exist, create it
		if err := db.Create(&employee).Error; err != nil {
			log.Println("Failed to create employee:", err)
		} else {
			log.Println("Created employee:", employee.Username)
		}
	}
}
