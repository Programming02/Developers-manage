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
	_, err := p.db.Exec(`
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
	rows, err := p.db.QueryContext(ctx, query, id)
	if err != nil {
		return models.Task{}, err
	}

	if err := rows.Scan(&task.Id, task.Title, task.Description, task.StartAt, task.FinishAt, task.Status, task.StartedAt, task.FinishedAt, task.ProgrammerId, task.ProjectId); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (p programmerRepo) CreateCommit(ctx context.Context, c models.Commit) error {
	return nil
}

func (p programmerRepo) UpdateCommit(ctx context.Context, c models.Commit) error {
	return nil
}

func (p programmerRepo) DeleteCommit(ctx context.Context, id string) error {
	return nil
}

func (p programmerRepo) GetCommitList(ctx context.Context) ([]models.Commit, error) {
	return []models.Commit{}, nil
}
