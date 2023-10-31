package models

type DriverDocuments struct {
	DocumentID   int    `json:"document_id" gorm:"primaryKey;unique"`
	UserID       int    `json:"user_id"`
	LicenseNo    string `json:"license_no" gorm:"not null"`
	LicenceExp   string `json:"license_exp" gorm:"not null"`
	LicenceFront string `json:"licence_ft_img" gorm:"not null"`
	LicenceBack  string `json:"licence_bk_img" gorm:"not null"`

	AdharNo      string `json:"adhar_no" gorm:"not null"`
	AdharAddress string `json:"ahdhar_address" gorm:"not null"`
	AdharFront   string `json:"Adhar_ft_img" gorm:"not null"`
	AdharBack    string `json:"Adhar_bk_img" gorm:"not null"`
}
