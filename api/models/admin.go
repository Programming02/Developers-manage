package models

type Users struct {
	Id          string `json:"id,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Role        string `json:"role,omitempty"`
	BirthDay    string `json:"birth_day,omitempty"`
	PhoneNumber string `json:"phone,omitempty"`
	Positions   string `json:"position,omitempty"`
}

type ListUsers []Users

type Project struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Status      string `json:"status,omitempty"`
	TeamLeadId  string `json:"teamlead_id,omitempty"`
	Attachments string `json:"attachment,omitempty"`
}

type ListProjects []Project

type Task struct {
	Id           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	StartAt      string `json:"start_at,omitempty"`
	FinishAt     string `json:"finish_at,omitempty"`
	Status       string `json:"status,omitempty"`
	StartedAt    string `json:"started_at,omitempty"`
	FinishedAt   string `json:"finished_at,omitempty"`
	ProgrammerId string `json:"programmer_id,omitempty"`
	Attachments  string `json:"attachments,omitempty"`
	ProjectId    string `json:"project_id,omitempty"`
}

type CheckTeamLeadRequest struct {
	UserId    string `json:"userId,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
}
