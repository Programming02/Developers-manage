package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/jwt"
	"github.com/programming02/osg/storage"
	"net/http"
)

type ProgrammerService struct {
	storage storage.IStorage
}

func NewProgrammerService(db *sqlx.DB) *ProgrammerService {
	return &ProgrammerService{
		storage: storage.New(db),
	}
}

func (p ProgrammerService) CreateTask(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	role, err := p.UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: task.ProjectId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	if role != "team_lead" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "You're not team_lead"})
	}

	if err := p.storage.Programmer().CreateTask(context.Background(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (p ProgrammerService) UpdateTask(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}
	task := models.Task{}
	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	role, err := p.UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: task.ProjectId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	if role != "team_lead" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "You're not team_lead"})
	}

	err = p.storage.Programmer().UpdateTask(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func (p ProgrammerService) DeleteTask(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errors.New("there's a mistake on 'id'"),
		})
	}

	task, err := p.storage.Programmer().GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	role, err := p.UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: task.ProjectId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	if role != "team_lead" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "You're not team_lead"})
	}

	err = p.storage.Programmer().DeleteTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (p ProgrammerService) GetTask(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errors.New("there's a mistake on 'id'"),
		})
	}

	t, err := p.storage.Programmer().GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	_, err = p.UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: t.ProjectId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"err": errors.New("you don't have access")})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	task, err := p.storage.Programmer().GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"ok": task})
	//c.JSON(http.StatusOK, gin.H{
	//	"id":            t.Id,
	//	"title":         t.Title,
	//	"description":   t.Description,
	//	"start_at":      t.StartAt,
	//	"finish_at":     t.FinishedAt,
	//	"status":        t.Status,
	//	"started_at":    t.StartedAt,
	//	"finished_at":   t.FinishedAt,
	//	"programmer_id": t.ProgrammerId,
	//	"attachments":   t.Attachments,
	//	"project_id":    t.ProjectId,
	//})
}

func (p ProgrammerService) CreateCommit(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	commit := models.Commit{}
	if err := c.ShouldBindJSON(commit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
	}

	t, err := p.storage.Programmer().GetTask(context.Background(), commit.TaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	_, err = p.UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: t.ProjectId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"err": errors.New("you don't have access")})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"err": err.Error(),
	})

	err = p.storage.Programmer().CreateCommit(context.Background(), commit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func (p ProgrammerService) UpdateCommit(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	com := models.Commit{}
	if err := c.ShouldBindJSON(com); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	err = p.storage.Programmer().UpdateCommit(context.Background(), com, user.UserID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (p ProgrammerService) GetCommits(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	taskID := c.Param("task_id")
	if err := c.ShouldBindJSON(taskID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": errors.New("there's mistake on 'id'")})
	}

	t, err := p.storage.Programmer().GetTask(context.Background(), taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	_, err = p.UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: t.ProjectId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"err": errors.New("you don't have access")})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"err": err.Error(),
	})

	res, err := p.storage.Programmer().GetCommitList(context.Background(), taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, res)
}

func (p ProgrammerService) DeleteCommit(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	err = p.storage.Programmer().DeleteCommit(context.Background(), user.UserID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (p ProgrammerService) CreateAttendance(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	types := c.Param("type")
	if types != "came" && types != "gone" {
		c.JSON(http.StatusBadRequest, gin.H{"err": errors.New("err with type")})
	}

	err = p.storage.Programmer().CreateAttendance(context.Background(), models.Attendance{
		UserId: user.UserID.String(),
		Type:   types,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	}
}

func (p ProgrammerService) UserRole(ctx context.Context, role models.UserRole) (string, error) {
	return p.storage.Programmer().UserRole(ctx, role)
}
