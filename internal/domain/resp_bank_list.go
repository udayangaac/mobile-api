package domain

type BankListResponse struct {
	Id           int    `json:"id"`
	BankName     string `json:"bank_name"`
	Image        string `json:"image"`
	IsSelected   int    `json:"is_selected"`
}
