package models

type LoginRequestModel struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
