package models

type Payment struct {
	ID        int     `json:"id" gorm:"primaryKey;unique"`
	BookingID int     `json:"booking_id" gorm:"not null"`
	User_id   int     `json:"userid"`
	Date      string  `json:"date" gorm:"not null"`
	DriverId  int     `json:"driver_id"`
	Distance  float64 `json:"distance"`
	Fare      int     `json:"fare"`
	Status    string  `json:"status"`
}
