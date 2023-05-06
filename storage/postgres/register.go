package postgres

import (
	"database/sql"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/storage/repo"
	"golang.org/x/crypto/bcrypt"
)

type registerRepo struct {
	db *sql.DB
}

func NewRegisterRepo(db *sql.DB) repo.Register {
	return &registerRepo{
		db: db,
	}
}

func (r *registerRepo) RegisterUser(u models.Users) error {
	_, err := r.db.Exec(`
	INSERT INTO users (id, full_name, avatar, role, birth_day, phone, position) VALUES ($1, $2, $3, $4, $5, $6, $7)
`,
		u.Id, u.FullName, u.Avatar, u.Role, u.BirthDay, u.PhoneNumber, u.Positions)
	if err != nil {
		return err
	}
	return nil
}

func (r *registerRepo) Login(req models.LoginRequestModel) error {
	var password string
	if err := r.db.QueryRow(`
		SELECT password FROM users WHERE id=$1
`,
		req.Phone).Scan(password); err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return err
	}
	return nil
}

//func (r *registerRepo) Login(req models.RegisterRequest) (res models.RegisterResponse, err error) {
//	if err = r.db.QueryRow(`
//	SELECT id, role FROM user
//`,
//		req.PhoneNumber, req.Password).Scan(&res.UserID, &res.Role); err != nil {
//		return res, err
//	}
//	return res, nil
//}
