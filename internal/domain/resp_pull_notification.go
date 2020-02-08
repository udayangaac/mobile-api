package domain

type PullResponse struct {
	ID                  	int    `json:"id"`
	NotificationCode    	string `json:"code"`
	ProviderName        	int    `json:"provider_name"`
	StatusLabel         	string `json:"status_label"`
	Status              	string `json:"status"`
	ProviderAvatar			string `json:"provider_avatar"`
	NotificationContent 	string `json:"notification_content"`
}