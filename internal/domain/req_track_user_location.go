package domain

type TrackUserLocation struct {
	UserId    int    `json:"userId"`
	Latitude  string `json:latitude`
	Longitude string `json:longitude`
}