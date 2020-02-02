package domain


type PullRequest struct {
	UserId    int16 `json:"id"`
	Type      string `json:"type"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}