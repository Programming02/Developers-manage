package models

type Project struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Status      string `json:"status,omitempty"`
	TeamLeadId  string `json:"teamlead_id,omitempty"`
	Attachments string `json:"attachment,omitempty"`
}

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

type UserRole struct {
	UserId    string `json:"userId"`
	ProjectId string `json:"project_id"`
}

type ListProjects []Project

type Commit struct {
	TaskID       string `json:"task_id,omitempty"`
	ProgrammerID string `json:"programmer_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
}

type Attendance struct {
	Type      string `json:"type,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
