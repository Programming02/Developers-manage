package postgres

import (
	"context"
	"database/sql"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/storage/repo"
)

type adminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) repo.Admin {
	return &adminRepo{
		db: db,
	}
}

func (a adminRepo) GetUser(ctx context.Context, id string) (models.Users, error) {
	query := `
	SELECT * FROM users WHERE id=$1
`
	user := models.Users{}
	err := a.db.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.FullName, &user.Avatar, &user.Role, &user.BirthDay, &user.PhoneNumber, &user.Positions)
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}

func (a adminRepo) CreateUser(ctx context.Context, d models.Users) error {
	_, err := a.db.Exec(`
	insert into users (id, full_name, avatar, role, birth_day, phone, position) values ($1, $2, $3, $4, $5, $6, $7) `,
		d.Id, d.FullName, d.Avatar, d.Role, d.BirthDay, d.PhoneNumber, d.Positions,
	)

	if err != nil {
		return err
	}

	return nil
}

func (a adminRepo) UpdateUser(ctx context.Context, u models.Users) error {
	_, err := a.db.ExecContext(ctx, `
	UPDATE users SET id=$1, full_name=$2, avatar=$3, role=$4, birth_day=$5, phone=$6, position=$7
`,
		u.Id, u.FullName, u.Avatar, u.Role, u.Role, u.BirthDay, u.PhoneNumber, u.Positions,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a adminRepo) DeleteUser(ctx context.Context, id string) error {
	query := `
	DELETE FROM users WHERE id=$1
`
	_, err := a.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (a adminRepo) CreateProject(ctx context.Context, d models.Project) error {
	_, err := a.db.Exec(`
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

func (a adminRepo) UpdateProject(ctx context.Context, p models.Project) error {
	_, err := a.db.ExecContext(ctx, `
	UPDATE project SET id=$1, name=$2, end_date=$3, status=$4, teamlead_id=$5, attachments=$6
`,
		p.Id, p.Name, p.EndDate, p.Status, p.TeamLeadId, p.Attachments,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a adminRepo) DeleteProject(ctx context.Context, id string) error {
	query := `
	DELETE FROM project WHERE id=$1
`
	_, err := a.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (a adminRepo) GetUserList(ctx context.Context) ([]models.Users, error) {
	query := `
	SELECT * FROM users
`
	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return []models.Users{}, err
	}

	res := models.ListUsers{}
	for rows.Next() {
		user := models.Users{}
		if err := rows.Scan(user.Id, user.FullName, user.Avatar, user.Role, user.BirthDay, user.PhoneNumber, user.Positions); err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	return res, nil
}

func (a adminRepo) GetProjectList(ctx context.Context) ([]models.Project, error) {
	query := `
	SELECT * FROM project
`
	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	res := models.ListProjects{}
	for rows.Next() {
		project := models.Project{}
		if err := rows.Scan(project.Id, project.Name, project.StartDate, project.EndDate, project.Status, project.Attachments); err != nil {
			return nil, err
		}
		res = append(res, project)
	}
	return res, nil
}

func (a adminRepo) GetProject(ctx context.Context, id string) (models.Project, error) {
	query := `
	SELECT * FROM project WHERE id=$1
`
	var project models.Project
	err := a.db.QueryRowContext(ctx, query, id).Scan(&project.Id, &project.Name, &project.StartDate, &project.EndDate, &project.Status, &project.TeamLeadId, &project.Attachments)

	if err != nil {
		return models.Project{}, err
	}

	return project, nil
}

//func (a adminRepo) CheckTeamLeadRequest(ctx context.Context, req models.CheckTeamLeadRequest) (bool, error) {
//	err := a.db.QueryRowContext(ctx, `
//	SELECT id FROM project WHERE teamlead_id = $1 AND id = $2
//`,
//		req.UserId, req.ProjectId).Scan(&req.UserId)
//	if errors.Is(err, errors.New("sql: no rows in result set")) {
//		return false, err
//	}
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}

//func (a adminRepo) GetUserRole(ctx context.Context, userId string) (string, error) {
//	var role string
//
//	err := a.db.QueryRow(`
//	SELECT role FROM users
//`,
//	userId).Scan(role)
//	if err != nil {
//		return "", err
//	}
//	return role, nil
//}
