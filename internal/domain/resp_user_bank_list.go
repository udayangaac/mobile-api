package domain


type UserBankListResponse struct {
	IsBankSelected int `json:"bank_selected"`
	BankList []BankListResponse `json:"bank_list"`
}