package models

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	Name         string
	Universities []University `gorm:"many2many:professor_universities;"`
	Departments  []Department `gorm:"many2many:professor_departments;"`
	Classes      []Class      `gorm:"many2many:professor_classes;"`
	Reviews      []Review
}
