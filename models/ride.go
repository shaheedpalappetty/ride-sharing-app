package models

type Ride struct {
	ID        int     `json:"id" gorm:"primaryKey;unique"`
	BookingID int     `json:"booking_id" gorm:"not null"`
	Date      string  `json:"date" gorm:"not null"`
	DriverId  int     `json:"driver_id"`
	Distance  float64 `json:"distance"`
	Fare      int     `json:"fare"`
}
