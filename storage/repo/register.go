package repo

import "github.com/programming02/osg/api/models"

type Register interface {
	Login(req models.RegisterRequest) (models.RegisterResponse, error)
}
