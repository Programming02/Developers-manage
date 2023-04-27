package repository

import (
	"context"
	"github.com/programming02/osg/api/models"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (models.Users, error)
	CreateUser(ctx context.Context, d models.Users) error
	UpdateUser(ctx context.Context, d models.Users) error
	DeleteUser(ctx context.Context, id string) error
	GetProject(ctx context.Context, id string) (models.Project, error)
	CreateProject(ctx context.Context, d models.Project) error
	UpdateProject(ctx context.Context, d models.Project) error
	DeleteProject(ctx context.Context, id string) error
	GetTask(ctx context.Context, id string) (models.Task, error)
	CreateTask(ctx context.Context, d models.Task) error
	UpdateTask(ctx context.Context, t models.Task) error
	DeleteTask(ctx context.Context, id string) error
	ProjectList(ctx context.Context) ([]models.Project, error)
	// TODO write comment
	// TODO programmer/attendance
}

type Register interface {
	Login(req models.LoginRequestModel)
}
