package repo

import (
	"context"
	"github.com/programming02/osg/api/models"
)

type Programmer interface {
	CreateTask(ctx context.Context, t models.Task) error
	CreateCommit(ctx context.Context, c models.Commit) error
	CreateAttendance(ctx context.Context, req models.Attendance) error
	GetTask(ctx context.Context, id string) (models.Task, error)
	GetCommit(ctx context.Context, taskId string) ([]models.Commit, error)
	UpdateTask(ctx context.Context, t models.Task) error
	UpdateCommit(ctx context.Context, c models.Commit, userID string) error
	DeleteTask(ctx context.Context, id string) error
	DeleteCommit(ctx context.Context, id string) error
	UserRole(ctx context.Context, role models.UserRole) (string, error)
}
