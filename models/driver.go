package models

type Driver struct {
	ID            int     `json:"id" gorm:"primaryKey;unique"`
	Name          string  `json:"name" gorm:"not null"`
	LastName      string  `json:"last_name" gorm:"not null"`
	PhoneNumber   string  `json:"phone_number" gorm:"not null"`
	Email         string  `json:"email" gorm:"not null;unique"`
	BirthDate     string  `json:"birth_date" gorm:"not null"`
	DriverImg     string  `json:"driver_img" gorm:"not null"`
	Gender        string  `json:"gender" gorm:"not null"`
	UserName      string  `json:"username"`
	Password      string  `json:"password"`
	LicenseNumber string  `json:"license_number" gorm:"not null"`
	Experience    string  `json:"experience" gorm:"not null"`
	Status        string  `json:"status"`
	Description   string  `json:"description"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Online        bool    `json:"online" gorm:"default:false"`

}
