package moduls

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	UserId      uuid.UUID     `json:"id,omitempty"`
	FullName    string        `json:"full_name,omitempty"`
	Avatar      string        `json:"avatar,omitempty"`
	Role        string        `json:"role,omitempty"`
	BirthDay    time.Duration `json:"birth_day,omitempty"`
	PhoneNumber string        `json:"phone,omitempty"`
	Positions   string        `json:"position,omitempty"`
}

type Project struct {
	Id          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	StartDate   time.Weekday `json:"start_date,omitempty"`
	EndDate     time.Weekday `json:"end_date,omitempty"`
	Status      string       `json:"status,omitempty"`
	TeamLeadId  uuid.UUID    `json:"teamlead_id,omitempty"`
	Attachments string       `json:"attachment,omitempty"`
}

type Task struct {
	Id           string    `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	StartAt      time.Time `json:"start_at,omitempty"`
	FinishAt     time.Time `json:"finish_at,omitempty"`
	Status       string    `json:"status,omitempty"`
	StartedAt    time.Time `json:"started_at,omitempty"`
	FinishedAt   time.Time `json:"finished_at,omitempty"`
	ProgrammerId string    `json:"programmer_id,omitempty"`
	Attachments  string    `json:"attachments,omitempty"`
	ProjectId    string    `json:"project_id,omitempty"`
}
