package models

type User struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	User_ID uint   `json:"user_id" gorm:"unique;not null;default: null"`
	Email   string `json:"email" gorm:"unique;not null;default: null"`
	Fname   string `json:"first_name" gorm:"not null;default: null"`
	Lname   string `json:"last_name" gorm:"not null;default: null"`
	Phone   string `json:"phone" gorm:"unique;not null;default: null"`
	Role    string `json:"role" gorm:"not null;default: null"`
}
