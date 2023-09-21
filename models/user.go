package models

type User struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	User_ID uint   `json:"id_user" gorm:"unique;not null;default: null"`
	Email   string `json:"email" gorm:"unique;not null;default: null"`
	Fname   string `json:"firstname" gorm:"not null;default: null"`
	Lname   string `json:"lastname" gorm:"not null;default: null"`
	Phone   string `json:"phone" gorm:"unique;not null;default: null"`
	Role    string `json:"role" gorm:"not null;default: null"`
}