package entities

type Notification struct {
	Id           int16 // `json:"name"`
	Content         string // `json:"email"`
	Status         int16  // `json:"status"`
	Address        string // json:"address"`
}