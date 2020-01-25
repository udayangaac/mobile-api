package domain


type PullRequest struct {
	UserId    string `json:"id"`
	Type      string `json:"type"`
	Location struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"location"`
}