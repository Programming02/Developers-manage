package postgres

import (
	"context"
	"database/sql"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/storage/repo"
)

type programmerRepo struct {
	db *sql.DB
}

func NewProgrammerRepo(db *sql.DB) repo.Programmer {
	return &programmerRepo{
		db: db,
	}
}

func (p programmerRepo) CreateTask(ctx context.Context, t models.CreateTaskRequestModel) error {
	query := `
INSERT INTO task (title, description, start_at, finish_at, programmer_id, attachments, project_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`
	if err := p.db.QueryRowContext(ctx, query, t.Title, t.Description, t.StartAt, t.FinishAt, t.ProgrammerID, t.Attachment, t.ProjectID); err != nil {
		return err.Err()
	}
	return nil
}

func (p programmerRepo) UpdateTask(ctx context.Context, t models.Task) error {
	_, err := p.db.ExecContext(ctx, `
	UPDATE task SET id=$1, title=$2, description=$3, finish_at=$4, status=$5, started_at=$6, finished_at=$7, programmer_id=$8, attachments=$9, project_id=$10
`,
		t.Id, t.Title, t.Description, t.FinishAt, t.Status, t.StartedAt, t.FinishedAt, t.ProgrammerId, t.ProjectId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (p programmerRepo) DeleteTask(ctx context.Context, id string) error {
	query := `
	DELETE FROM task WHERE id=$1
`
	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p programmerRepo) GetTask(ctx context.Context, id string) (models.Task, error) {
	query := `
	SELECT * FROM task WHERE id=$1
`
	task := models.Task{}
	err := p.db.QueryRowContext(ctx, query, id).Scan(&task.Id, task.Title, task.Description, task.StartAt, task.FinishAt, task.Status, task.StartedAt, task.FinishedAt, task.ProgrammerId, task.ProjectId)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (p programmerRepo) CreateCommit(ctx context.Context, c models.Commit) error {
	query := `
	INSERT INTO comments VALUES ($1, $2, $3)
`
	_, err := p.db.ExecContext(ctx, query, c.TaskID, c.ProgrammerID, c.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p programmerRepo) UpdateCommit(ctx context.Context, c models.Commit, userID string) error {
	query := `
	UPDATE comments SET task_id=$1, user_id=$2, created_at=$3 WHERE id=$4
`
	_, err := p.db.ExecContext(ctx, query, c.TaskID, c.ProgrammerID, c.CreatedAt, userID)
	if err != nil {
		return err
	}
	return nil
}

func (p programmerRepo) DeleteCommit(ctx context.Context, id string) error {
	query := `
	DELETE FROM comments WHERE id=$1
`
	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p programmerRepo) GetCommitList(ctx context.Context, taskID string) ([]models.Commit, error) {
	rows, err := p.db.QueryContext(ctx, `
	SELECT * FROM comments WHERE id=$1
`,
		taskID)
	if err != nil {
		return []models.Commit{}, err
	}
	var res []models.Commit
	for rows.Next() {
		com := models.Commit{}
		err := rows.Scan(&com.TaskID, &com.ProgrammerID, &com.CreatedAt)
		if err != nil {
			return []models.Commit{}, err
		}
		res = append(res, com)
	}
	return res, nil
}

func (p programmerRepo) CreateAttendance(ctx context.Context, req models.Attendance) error {
	_, err := p.db.ExecContext(ctx, `
	INSERT INTO attendance VALUES ($1, $2, $3)
`,
		req.Type, req.UserId, req.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (p programmerRepo) GetAttendanceList(ctx context.Context, req models.GetUserAttendanceRequest) ([]models.Attendance, error) {
	query := `
	SELECT * FROM attendance WHERE user_id=$1 AND type=$2
`
	rows, err := p.db.QueryContext(ctx, query, req.UserID, req.Type)
	if err != nil {
		return []models.Attendance{}, err
	}
	defer rows.Close()
	var attendanceList []models.Attendance
	for rows.Next() {
		var attendance models.Attendance
		if err := rows.Scan(&attendance.Type, &attendance.UserId, &attendance.CreatedAt); err != nil {
			return nil, err
		}
		attendanceList = append(attendanceList, attendance)
	}
	return attendanceList, nil
}
