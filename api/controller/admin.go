package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/moduls"
	// "go get -u github.com/swaggo/files"
	// "go get -u github.com/swaggo/gin-swagger"
	"net/http"
)

func (a Api) GetAdmin(c *gin.Context) {
	id := c.Param("id")
	b, err := a.Repo.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id":   b.UserId,
		"full_name": b.FullName,
		"avatar":    b.Avatar,
		"Role":      b.Role,
		"birth_day": b.BirthDay,
		"phone":     b.PhoneNumber,
		"position":  b.Positions,
	})
}

func (a Api) CreateAdmin(c *gin.Context) {
	admin := moduls.Users{}
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := a.Repo.CreateUser(context.Background(), admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

/* func (a Api) UpdateAdmin(c *gin.Context) {
	admin := moduls.Users{}
	if err := c.ShouldBindJSON(admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if err := a.Repo.UpdateUser(context.Background(), admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}
*/

func (a Api) DeleteAdmin(c *gin.Context) {
	id := c.Query("id")

	err := a.Repo.DeleteUser(c.Request.Context(), id)
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

func (a Api) GetProject(c *gin.Context) {
	id := c.Param("id")
	b, err := a.Repo.GetProject(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          b.Id,
		"name":        b.Name,
		"start_date":  b.StartDate,
		"end_date":    b.EndDate,
		"status":      b.Status,
		"teamlead_id": b.TeamLeadId,
		"attachment":  b.Attachments,
	})
}

func (a Api) CreateProject(c *gin.Context) {
	project := moduls.Project{}
	if err := c.ShouldBindJSON(project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if err := a.Repo.CreateProject(context.Background(), project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (a Api) DeleteProject(c *gin.Context) {
	name := c.Query("name")

	err := a.Repo.DeleteProject(context.Background(), name)
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

func (a Api) GetTask(c *gin.Context) {
	id := c.Param("id")
	t, err := a.Repo.GetTask(context.Background(), id)
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

func (a Api) CreateTask(c *gin.Context) {
	task := moduls.Task{}

	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	if err := a.Repo.CreateTask(context.Background(), task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
