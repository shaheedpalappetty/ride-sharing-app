package models

type Coupons struct {
	Code       string  `json:"code"`
	Percentage float32 `json:"percentage"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
}
