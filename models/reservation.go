package models

type Reservation struct {
	Id          int    `json:"id_room" gorm:"primarykey"`
	Rname       string `json:"roomname"`
	Instructor  string `json:"instructor"`
	Phone       string `json:"phone"`
	Description string `json:"Description"`
	Date        string `json:"date"`
	TimeS       string `json:"time-start"`
	TimeE       string `json:"time-end"`
	Status      string `json:"status"`
}
