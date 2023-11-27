package models

type Firebase struct {
	ID       int    `json:"id" gorm:"primaryKey;unique"`
	UserId   int    `json:"user_id" gorm:"not null"`
	Category string `json:"category" gorm:"not null"`
	Token    string `json:"token" gorm:"not null"`
	Status   string `json:"status" gorm:"not null"`
}
