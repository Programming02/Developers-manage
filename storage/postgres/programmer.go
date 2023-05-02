package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/storage/repo"
)

type programmerRepo struct {
	db *sqlx.DB
}

func NewProgrammerRepo(db *sqlx.DB) repo.Programmer {
	return &programmerRepo{
		db: db,
	}
}

func (p programmerRepo) CreateTask(ctx context.Context, t models.Task) error {
	_, err := p.db.ExecContext(ctx, `
	INSERT INTO task (id, title, description, finish_at, status, started_at, finished_at, programmer_id, attachments, project_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
`,
		t.Id, t.Title, t.Description, t.FinishAt, t.Status, t.StartedAt, t.FinishedAt, t.ProgrammerId, t.Attachments, t.ProjectId,
	)
	if err != nil {
		return err
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
	err := p.db.QueryRowContext(ctx, query, task).Scan(&task.Id, task.Title, task.Description, task.StartAt, task.FinishAt, task.Status, task.StartedAt, task.FinishedAt, task.ProgrammerId, task.ProjectId)
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

func (p programmerRepo) GetCommitList(ctx context.Context, taskId string) ([]models.Commit, error) {
	rows, err := p.db.QueryContext(ctx, `
	SELECT * FROM commit WHERE id=$1
`,
		taskId)
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

func (p programmerRepo) UserRole(ctx context.Context, role models.UserRole) (string, error) {
	return "", nil
}
