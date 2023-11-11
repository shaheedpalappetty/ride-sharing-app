package models

type Coupons struct {
	Code       string `json:"code"`
	Percentage string `json:"percentage"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}
