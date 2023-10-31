package models

type VehicleDetails struct {
	DetilsID      int    `json:"detils_id" gorm:"primaryKey;unique"`
	UserID        int    `json:"user_id"`
	VehicleBrand  string `json:"vehicle_brand" gorm:"not null"`
	VehicleModel  string `json:"vehicle_model" gorm:"not null"`
	VehicleYear   string `json:"vehicle_year" gorm:"not null"`
	VehicleColor  string `json:"vehicle_color" gorm:"not null"`
	VehicleSeat   string `json:"vehicle_seat" gorm:"not null"`
	VehicleNumber string `json:"vehicle_number" grom:"not null"`
}
