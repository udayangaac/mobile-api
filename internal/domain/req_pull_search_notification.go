package domain

type PullSearchRequest struct {
	UserId      int 	 `json:"userId"`
    SearchText  string   `json:"searchText"`
}