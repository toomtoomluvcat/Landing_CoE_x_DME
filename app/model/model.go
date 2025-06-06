package model

import "gorm.io/gorm"

type Group struct {
	ID        uint
	Name      string
	Joins     []Join
	Code      string `gorm:"uniqueIndex"`
	ProblemID int64  `gorm:"foreignKey:GroupID"`
}

type Join struct {
	gorm.Model
	Email       string `gorm:"size:255;uniqueIndex:idx_email"`
	Profile     string
	UserName    string
	GroupID     *uint
	Description string
	Link_1      string
	Link_2      string
	Link_3      string
}
