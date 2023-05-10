package models

type LoginRequestModel struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type GetUserAttendanceRequest struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}

type CreateTaskRequestModel struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	StartAt      string `json:"start_at"`
	FinishAt     string `json:"finish_at"`
	ProgrammerID string `json:"programmer_id"`
	Attachment   string `json:"attachment"`
	ProjectID    int    `json:"project_id"`
}

type DeleteCommentRequest struct {
	ID       string `json:"ID"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type GetCommentsRequestModel struct {
	TaskID int `json:"task_id"`
}
