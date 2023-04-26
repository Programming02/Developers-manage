package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/programming02/osg/api/moduls"
)

type Server struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Server {
	return Server{
		db: db,
	}
}

func (s Server) GetUser(ctx context.Context, id string) (moduls.Users, error) {
	query := `
	SELECT * FROM users WHERE id=$1
`
	var user moduls.Users
	err := s.db.QueryRow(query, id).Scan(&user.Id, &user.FullName, &user.Avatar, &user.Role, &user.BirthDay, &user.PhoneNumber, &user.Positions)
	if err != nil {
		fmt.Println(err.Error())
		return moduls.Users{}, err
	}

	return user, nil
}

func (s Server) CreateUser(ctx context.Context, d moduls.Users) error {
	_, err := s.db.Exec(`
	insert into users (id, full_name, avatar, role, birth_day, phone, position) values ($1, $2, $3, $4, $5, $6, $7) `,
		d.Id, d.FullName, d.Avatar, d.Role, d.BirthDay, d.PhoneNumber, d.Positions,
	)

	if err != nil {
		return err
	}

	return nil
}

//func (s Server) UpdateUser(c *gin.Context) {
//	query := `
//	UPDATE
//`
//}

func (s Server) DeleteUser(ctx context.Context, id string) error {
	query := `
	DELETE FROM users WHERE id=$1
`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Server) GetProject(ctx context.Context, id string) (moduls.Project, error) {
	query := `
	SELECT * FROM project WHERE id=$1
`
	var project moduls.Project
	err := s.db.QueryRowContext(ctx, query, id).Scan(&project.Id, &project.Name, &project.StartDate, &project.EndDate, &project.Status, &project.TeamLeadId, &project.Attachments)

	if err != nil {
		return moduls.Project{}, err
	}

	return project, nil
}

func (s Server) CreateProject(ctx context.Context, d moduls.Project) error {
	_, err := s.db.Exec(`
	INSERT INTO project (id, name, end_date, status, teamlead_id, attachments)
	VALUES ($1, $2, $3, $4, $5, $6)
`,
		d.Id, d.Name, d.EndDate, d.Status, d.TeamLeadId, d.Attachments,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s Server) DeleteProject(ctx context.Context, id string) error {
	query := `
	DELETE FROM project WHERE id=$1
`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Server) GetTask(ctx context.Context, id string) (moduls.Task, error) {
	query := `
	SELECT * FROM task WHERE id=$1
`
	task := moduls.Task{}
	rows, err := s.db.QueryContext(ctx, query, id)
	if err != nil {
		return moduls.Task{}, err
	}

	if err := rows.Scan(&task.Id, task.Title, task.Description, task.StartAt, task.FinishAt, task.Status, task.StartedAt, task.FinishedAt, task.ProgrammerId, task.ProjectId); err != nil {
		return moduls.Task{}, err
	}
	return task, nil
}

func (s Server) CreateTask(ctx context.Context, t moduls.Task) error {
	_, err := s.db.Exec(`
	INSERT INTO task VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`,
		t.Id, t.Title, t.Description, t.StartAt, t.FinishAt, t.Status, t.StartedAt, t.FinishedAt, t.ProgrammerId, t.Attachments, t.ProjectId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s Server) DeleteTask(ctx context.Context, id string) error {
	query := `
	DELETE FROM task WHERE id=$1
`
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
