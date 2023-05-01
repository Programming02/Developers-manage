package repo

import (
	"context"
	"github.com/programming02/osg/api/models"
)

type Programmer interface {
	CreateTask(ctx context.Context, t models.Task) error
	UpdateTask(ctx context.Context, t models.Task) error
	DeleteTask(ctx context.Context, id string) error
	GetTask(ctx context.Context, id string) (models.Task, error)
	CreateCommit(ctx context.Context, c models.Commit) error
	UpdateCommit(ctx context.Context, c models.Commit, userID string) error
	DeleteCommit(ctx context.Context, id string) error
	GetCommitList(ctx context.Context, taskId string) ([]models.Commit, error)
	UserRole(ctx context.Context, role models.UserRole) (string, error)
	CreateAttendance(ctx context.Context, req models.Attendance) error
}
