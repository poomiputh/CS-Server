package models

type User struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	CollegeID          uint      `json:"college_id" gorm:"unique;default: null"`
	Password         string    `json:"password" gorm:"not null;default: null"`
	Email            string    `json:"email" gorm:"unique;not null;default: null"`
	Fname            string    `json:"first_name" gorm:"not null;default: null"`
	Lname            string    `json:"last_name" gorm:"default: null"`
	Phone            string    `json:"phone" gorm:"unique;not null;default: null"`
	Role             string    `json:"role" gorm:"not null;default: null"`
	Request          []Request `json:"request" gorm:"foreignKey:UserRefer"`
	ApprovedRequest []Request `json:"approved_request" gorm:"foreignKey:AdminRefer"`
}
