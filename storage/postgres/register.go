package postgres

import (
	"database/sql"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/storage/repo"
)

type registerRepo struct {
	db *sql.DB
}

func NewRegisterRepo(db *sql.DB) repo.Register {
	return &registerRepo{
		db: db,
	}
}

func (r *registerRepo) Login(req models.RegisterRequest) (res models.RegisterResponse, err error) {
	if err = r.db.QueryRow(`
	SELECT id, role FROM user
`,
		req.PhoneNumber, req.Password).Scan(&res.UserID, &res.Role); err != nil {
		return res, err
	}
	return res, nil
}
