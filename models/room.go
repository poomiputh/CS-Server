package models

type Room struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Rname        string `json:"roomname" gorm:"unique;not null;default: null"`
	Rstatus      bool   `json:"roomstatus" gorm:"not null;default: null"`
}
