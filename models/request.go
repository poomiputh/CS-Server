package models

type Request struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Room_ID       string `json:"room_id" gorm:"not null;default: null"`
	UserRefer     uint   `json:"user_refer" gorm:"not null;default: null"`
	Data_User     []User `json:"data_user" gorm:"foreignKey:ID;references:UserRefer"`
	AdminRefer    uint   `json:"admin_refer" gorm:"default: null"`
	Data_Admin    []User `json:"data_admin" gorm:"foreignKey:ID;references:AdminRefer"`
	Instructor    string `json:"instructor" gorm:"default: null"`
	ReDescription string `json:"request_description" gorm:"default: null"`
	StartTime     string `json:"start_time" gorm:"not null;default: null"`
	EndTime       string `json:"end_time" gorm:"not null;default: null"`
	Date          string `json:"date" gorm:"not null;default: null"`
	Status        string `json:"status" gorm:"not null;default: null"`
}
