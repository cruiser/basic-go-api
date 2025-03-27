package models

type Client struct {
	ClientId             uint    `gorm:"primaryKey" json:"clientId"`
	Name                 string  `json:"name"`
	PhoneNumber          string  `json:"phoneNumber"`
	Email                string  `json:"email"`
	Spent                float64 `json:"spent"`
	LastChoreographyDate string  `json:"lastChoreographyDate"`
	Status               string  `json:"status"`
	Notes                string  `json:"notes"`
}
