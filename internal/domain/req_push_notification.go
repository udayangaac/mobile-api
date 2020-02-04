package domain

type PushRequest struct {
	UserId    int `json:"id"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}