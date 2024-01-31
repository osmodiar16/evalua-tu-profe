package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Content         string
	Rating          uint
	Difficulty      uint
	Textbook        bool
	Attendance      bool
	WouldTakeAgain  bool
	Grade           uint
	ThumbsDownTotal uint
	ThumbsUpTotal   uint
	ProfessorId     uint
	UserId          uint
}
