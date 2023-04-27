package models

type LoginRequestModel struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
