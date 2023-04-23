package repository

import (
	"context"
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
	SELECT * FROM users WHERE id=$1,
`
	var user moduls.Users
	err := s.db.QueryRowContext(ctx, query, id).Scan(&user.UserId, &user.FullName, &user.Avatar, &user.Role, &user.BirthDay, &user.PhoneNumber, &user.Positions)
	if err != nil {
		return moduls.Users{}, err
	}

	return user, nil
}

func (s Server) CreateUser(ctx context.Context, d moduls.Users) error {
	_, err := s.db.Exec(`
	insert into users (id, full_name, avatar, role, birth_day, phone, position) values ($1, $2, $3, $4, $5, $6, $7) `,
		d.UserId, d.FullName, d.Avatar, d.Role, d.BirthDay, d.PhoneNumber, d.Positions,
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

func (s Server) CreateProject(ctx context.Context, d moduls.Project) error {
	_, err := s.db.Exec(`
	INSERT INTO project VALUES ($1, $2, $3, $4, $5, $6)`,
		d.Name, d.StartDate, d.EndDate, d.Status, d.TeamLeadId, d.Attachments,
	)
	if err != nil {
		return err
	}
	return nil
}
