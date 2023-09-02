package models

type Room struct {
	Id      int    `json:"id_room" gorm:"primarykey"`
	Rname   string `json:"roomname"`
	Rstatus bool   `json:"roomstatus"`
}
