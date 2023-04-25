package repository

import (
	"context"
	"github.com/programming02/osg/api/moduls"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (moduls.Users, error)
	CreateUser(ctx context.Context, d moduls.Users) error
	//UpdateUser(ctx context.Context, d moduls.Users) error
	DeleteUser(ctx context.Context, id string) error
	GetProject(ctx context.Context, id string) (moduls.Project, error)
	CreateProject(ctx context.Context, d moduls.Project) error
	//UpdateProject(ctx context.Context, d moduls.Project) error
	DeleteProject(ctx context.Context, name string) error
	GetTask(ctx context.Context, id string) (moduls.Task, error)
	CreateTask(ctx context.Context, d moduls.Task) error

	// TODO task crud
	// TODO programmer/project_list
	// TODO write comment
	// TODO programmer/attendance
}
