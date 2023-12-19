package models


type Coordinate struct {
	SLatitude  float64 `json:"slatitude"`
	SLongitude float64 `json:"slongitude"`
	ELatitude  float64 `json:"elatitude"`
	ELongitude float64 `json:"elongitude"`
}
type Location struct {
	ID        int
	Latitude  float64
	Longitude float64
}