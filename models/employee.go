package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name        string     `gorm:"size:255;not null;unique"`
	Description string     `gorm:"size:1024"`
	Employees   []Employee `gorm:"foreignKey:DepartmentID"`
}

type Employee struct {
	gorm.Model
	FirstName    string `gorm:"size:255;not null"`
	LastName     string `gorm:"size:255;not null"`
	Username     string `gorm:"size:255;not null;unique"`
	Password     string `gorm:"size:255;not null"`
	Email        string `gorm:"size:255;not null;unique"`
	DOB          time.Time
	DepartmentID uint
	Position     string     `gorm:"size:255;not null"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
}
