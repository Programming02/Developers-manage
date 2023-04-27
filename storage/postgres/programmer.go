package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/programming02/osg/storage/repo"
)

type programmerRepo struct {
	db *sqlx.DB
}

func NewProgrammerRepo(db *sqlx.DB) repo.Programmer {
	return &programmerRepo{db: db}
}
