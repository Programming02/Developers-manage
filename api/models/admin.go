package models

type Users struct {
	Id          string `json:"id,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Password    string `json:"password,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Role        string `json:"role,omitempty"`
	BirthDay    string `json:"birth_day,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`
	Positions   string `json:"position,omitempty"`
}

type ListUsers []Users

type CheckTeamLeadRequest struct {
	UserId    string `json:"userId,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
}
