package repo

import "github.com/programming02/osg/api/models"

type Register interface {
	RegisterUser(user models.Users) error
	Login(req models.LoginRequestModel) error
}
