package models

type User struct {
	Id      int    `json:"id" gorm:"primarykey"`
	ID_user int    `json:"id_user"`
	Fname   string `json:"firstname"`
	Lname   string `json:"lastname"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Role    string `json:"role"`
}
