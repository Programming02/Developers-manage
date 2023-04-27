package repo

import (
	"context"
	"github.com/programming02/osg/api/models"
)

type Programmer interface {
	CreateTask(ctx context.Context, task models.Task) error
	UpdateTask(ctx context.Context, task models.Task) error
	DeleteTask(ctx context.Context, id string) error
	GetTask(ctx context.Context, id string) (models.Task, error)
}
