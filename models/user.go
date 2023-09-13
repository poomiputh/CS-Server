package models

type User struct {
	ID             uint        `json:"id_user" gorm:"primaryKey"`
	UserRefer      uint        `json:"id" gorm:"unique;not null;default: null"`
	Data_GuestUser []Data_User `json:"data_admin" gorm:"foreignKey:ID;references:UserRefer"`
}
