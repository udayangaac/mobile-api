package domain

type PullRequest struct {
	UserId    int `json:"userId"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}