package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/models"
	"net/http"
)

func (a AdminService) CreateTask(c *gin.Context) {
	task := models.Task{}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	if err := a.storage.Programmer().CreateTask(context.Background(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (a AdminService) UpdateTask(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	err := a.storage.Programmer().UpdateTask(context.Background(), task)
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

func (a AdminService) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := a.storage.Programmer().DeleteTask(context.Background(), id)
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

func (a AdminService) GetTask(c *gin.Context) {
	id := c.Param("id")
	t, err := a.storage.Programmer().GetTask(context.Background(), id)
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
