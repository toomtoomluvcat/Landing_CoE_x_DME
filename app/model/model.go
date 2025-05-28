package model

import "gorm.io/gorm"

type Participant struct {
	Email       string
	Name        string
	UserName    string
	StudentID   int64
	GroupID     uint
	Role        string
	Description string
	Link_1      string
	Link_2      string
	Link_3      string
}

type Group struct {
	ID           uint
	Participants []Participant
}
type Join struct {
	gorm.Model
	Email   string `gorm:"size:255;uniqueIndex:idx_email"`
	Profile string
}
