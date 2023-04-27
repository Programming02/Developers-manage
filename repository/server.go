package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/programming02/osg/api/models"
)

type Server struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Server {
	return Server{
		db: db,
	}
}

func (s Server) GetUser(ctx context.Context, id string) (models.Users, error) {
	query := `
	SELECT * FROM users WHERE id=$1
`
	var user models.Users
	err := s.db.QueryRow(query, id).Scan(&user.Id, &user.FullName, &user.Avatar, &user.Role, &user.BirthDay, &user.PhoneNumber, &user.Positions)
	if err != nil {
		fmt.Println(err.Error())
		return models.Users{}, err
	}

	return user, nil
}

//func (s Server) CreateUser(ctx context.Context, d models.Users) error {
//	_, err := s.db.Exec(`
//	insert into users (id, full_name, avatar, role, birth_day, phone, position) values ($1, $2, $3, $4, $5, $6, $7) `,
//		d.Id, d.FullName, d.Avatar, d.Role, d.BirthDay, d.PhoneNumber, d.Positions,
//	)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//func (s Server) UpdateUser(ctx context.Context, u models.Users) error {
//	_, err := s.db.ExecContext(ctx, `
//	UPDATE user SET id=$1, full_name=$2, avatar=$3, role=$4, birth_day=$5, phone=$6, position=$7
//`,
//		u.Id, u.FullName, u.Avatar, u.Role, u.Role, u.BirthDay, u.PhoneNumber, u.Positions,
//	)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (s Server) DeleteUser(ctx context.Context, id string) error {
//	query := `
//	DELETE FROM users WHERE id=$1
//`
//	_, err := s.db.ExecContext(ctx, query, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (s Server) GetProject(ctx context.Context, id string) (models.Project, error) {
//	query := `
//	SELECT * FROM project WHERE id=$1
//`
//	var project models.Project
//	err := s.db.QueryRowContext(ctx, query, id).Scan(&project.Id, &project.Name, &project.StartDate, &project.EndDate, &project.Status, &project.TeamLeadId, &project.Attachments)
//
//	if err != nil {
//		return models.Project{}, err
//	}
//
//	return project, nil
//}

//func (s Server) CreateProject(ctx context.Context, d models.Project) error {
//	_, err := s.db.Exec(`
//	INSERT INTO project (id, name, end_date, status, teamlead_id, attachments)
//	VALUES ($1, $2, $3, $4, $5, $6)
//`,
//		d.Id, d.Name, d.EndDate, d.Status, d.TeamLeadId, d.Attachments,
//	)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s Server) UpdateProject(ctx context.Context, p models.Project) error {
//	_, err := s.db.ExecContext(ctx, `
//	UPDATE project SET id=$1, name=$2, end_date=$3, status=$4, teamlead_id=$5, attachments=$6
//`,
//		p.Id, p.Name, p.EndDate, p.Status, p.TeamLeadId, p.Attachments,
//	)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s Server) DeleteProject(ctx context.Context, id string) error {
//	query := `
//	DELETE FROM project WHERE id=$1
//`
//	_, err := s.db.ExecContext(ctx, query, id)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (s Server) GetTask(ctx context.Context, id string) (models.Task, error) {
	query := `
	SELECT * FROM task WHERE id=$1
`
	task := models.Task{}
	rows, err := s.db.QueryContext(ctx, query, id)
	if err != nil {
		return models.Task{}, err
	}

	if err := rows.Scan(&task.Id, task.Title, task.Description, task.StartAt, task.FinishAt, task.Status, task.StartedAt, task.FinishedAt, task.ProgrammerId, task.ProjectId); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s Server) CreateTask(ctx context.Context, t models.Task) error {
	_, err := s.db.Exec(`
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

func (s Server) UpdateTask(ctx context.Context, t models.Task) error {
	_, err := s.db.ExecContext(ctx, `
	UPDATE task SET id=$1, title=$2, description=$3, finish_at=$4, status=$5, started_at=$6, finished_at=$7, programmer_id=$8, attachments=$9, project_id=$10
`,
		t.Id, t.Title, t.Description, t.FinishAt, t.Status, t.StartedAt, t.FinishedAt, t.ProgrammerId, t.ProjectId,
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

//func (s Server) ProjectList(ctx context.Context) ([]models.Project, error) {
//	query := `
//	SELECT * FROM project
//`
//	rows, err := s.db.QueryContext(ctx, query)
//	if err != nil {
//		return nil, err
//	}
//	res := models.ListProjects{}
//	for rows.Next() {
//		project := models.Project{}
//		if err := rows.Scan(project.Id, project.Name, project.StartDate, project.EndDate, project.Status, project.Attachments); err != nil {
//			return nil, err
//		}
//		res = append(res, project)
//	}
//	return res, nil
//}
