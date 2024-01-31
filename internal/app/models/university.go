package models

import "gorm.io/gorm"

type University struct {
	gorm.Model
	Name        string
	Departments []Department
	Classes     []Class
}
