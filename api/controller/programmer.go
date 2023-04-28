package controller

import (
	"context"
	"github.com/gin-gonic/gin"
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
		return
	}
	task := models.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
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
	task := models.Task{}
	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	err := p.storage.Programmer().UpdateTask(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func (p ProgrammerService) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := p.storage.Programmer().DeleteTask(context.Background(), id)
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
	id := c.Param("id")
	t, err := p.storage.Programmer().GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            t.Id,
		"title":         t.Title,
		"description":   t.Description,
		"start_at":      t.StartAt,
		"finish_at":     t.FinishedAt,
		"status":        t.Status,
		"started_at":    t.StartedAt,
		"finished_at":   t.FinishedAt,
		"programmer_id": t.ProgrammerId,
		"attachments":   t.Attachments,
		"project_id":    t.ProjectId,
	})
}

func (p ProgrammerService) CreateCommit(c *gin.Context) {
	commit := models.Commit{}

	if err := c.ShouldBindJSON(commit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	err := p.storage.Programmer().CreateCommit(context.Background(), commit)
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


func (p ProgrammerService) UpdateCommit(c *gin.Context) {
	com := models.Commit{}

	if err := c.ShouldBindJSON(com); err != {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (p ProgrammerService) GetCommits(c *gin.Context) {

	res, err := p.storage.Programmer().GetCommitList(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (p ProgrammerService) DeleteCommit(c *gin.Context) {
	id := c.Param("id")

	err := p.storage.Programmer().DeleteCommit(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (p ProgrammerService) UserRole(c *gin.Context)  {

}