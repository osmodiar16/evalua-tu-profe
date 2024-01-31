package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name         string
	Code         string
	UniversityId uint
	DepartmentId uint
}
