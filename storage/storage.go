package storage

import (
	"database/sql"
	//"github.com/jmoiron/sqlx"
	"github.com/programming02/osg/storage/postgres"
	"github.com/programming02/osg/storage/repo"
)

type IStorage interface {
	Admin() repo.Admin
	Programmer() repo.Programmer
	Register() repo.Register
}

type storagePg struct {
	db         *sql.DB
	admin      repo.Admin
	programmer repo.Programmer
	register   repo.Register
}

func New(db *sql.DB) *storagePg {
	return &storagePg{
		db:         db,
		admin:      postgres.NewAdminRepo(db),
		programmer: postgres.NewProgrammerRepo(db),
		register:   postgres.NewRegisterRepo(db),
	}
}

func (s *storagePg) Admin() repo.Admin {
	return s.admin
}

func (s *storagePg) Programmer() repo.Programmer {
	return s.programmer
}

func (s *storagePg) Register() repo.Register {
	return s.register
}
