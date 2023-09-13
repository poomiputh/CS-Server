package models

type Admin struct {
	ID         uint        `json:"id_admin" gorm:"primaryKey"`
	UserRefer  uint        `json:"id" gorm:"unique;not null;default: null"`
	Data_Admin []Data_User `json:"data_admin" gorm:"foreignKey:ID;references:UserRefer"`
}
