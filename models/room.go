package models

type Room struct {
	ID               uint              `json:"id" gorm:"primaryKey"`
	Room_ID          uint              `json:"room_id" gorm:"not null;default: null"`
	Data_Reservation []ReservationTime `json:"data_reservationtime" gorm:"foreignKey:RoomRefer; references:Room_ID"`
}
