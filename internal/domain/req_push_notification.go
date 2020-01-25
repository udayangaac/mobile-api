package domain

type PushRequest struct {
	UserId    string `json:"id"`
	Location struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"location"`
}