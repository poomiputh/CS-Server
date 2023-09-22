package models

type Request struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Room_ID       string `json:"id_room" gorm:"not null;default: null"`
	UserRefer     uint   `json:"id_user" gorm:"not null;default: null"`
	Data_User     []User `json:"data_user" gorm:"foreignKey:ID;references:UserRefer"`
	AdminRefer    uint   `json:"id_admin" gorm:"default: null"`
	Data_Admin    []User `json:"data_admin" gorm:"foreignKey:ID;references:AdminRefer"`
	Instructor    string `json:"instructor" gorm:"default: null"`
	ReDescription string `json:"requestdescription" gorm:"default: null"`
	StartTime     string `json:"starttime" gorm:"not null;default: null"`
	EndTime       string `json:"endtime" gorm:"not null;default: null"`
	Date          string `json:"date" gorm:"not null;default: null"`
	Status        string `json:"status" gorm:"not null;default: null"`
}
