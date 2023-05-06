package repo

import (
	"context"
	"github.com/programming02/osg/api/models"
)

type Admin interface {
	GetUser(ctx context.Context, id string) (models.Users, error)
	GetUserList(ctx context.Context) ([]models.Users, error)
	GetProjectList(ctx context.Context) ([]models.Project, error)
	GetProject(ctx context.Context, id string) (models.Project, error)
	CreateUser(ctx context.Context, d models.Users) error
	CreateProject(ctx context.Context, d models.Project) error
	UpdateProject(ctx context.Context, d models.Project) error
	UpdateUser(ctx context.Context, d models.Users) error
	DeleteUser(ctx context.Context, id string) error
	DeleteProject(ctx context.Context, id string) error
}
