package models

type Room struct {
	RoomID          uint              `json:"room_id" gorm:"primaryKey"`
	DataReservation []ReservationTime `json:"data_reservationtime" gorm:"foreignKey:RoomRefer; references:Room_ID"`
}
