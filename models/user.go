package models

type User struct {
	ID          int    `json:"id" gorm:"primaryKey;unique"`
	FullName    string `json:"fullname" gorm:"not null"`
	BirthDate   string `json:"birthdate" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Password    string `json:"password" gorm:"not null"`
	PhoneNumber string `json:"phonenumber" gorm:"not null"`
	Gender      string `json:"gender" gorm:"not null"`
	Image       string `json:"image" gorm:"not null"`
}
