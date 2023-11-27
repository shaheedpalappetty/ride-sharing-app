package models

type Booking struct {
	ID          int     `json:"id" gorm:"primaryKey;unique"`
	UserID      int     `json:"userid" gorm:"not null"`
	PickupLat   float64 `json:"pickuplat" gorm:"not null"`
	PickupLong  float64 `json:"pickuplong" gorm:"not null"`
	DropoffLat  float64 `json:"droppofflat" gorm:"not null"`
	DropoffLong float64 `json:"dropofflong" gorm:"not null"`
	Status      string  `json:"status" gorm:"not null"`
}
