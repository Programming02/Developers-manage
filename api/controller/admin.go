package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/models"
	"time"

	"net/http"
)

func (a Api) GetAdmin(c *gin.Context) {
	id := c.Param("id")
	b, err := a.storage.Admin().GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user_id":   b.Id,
		"full_name": b.FullName,
		"password":  b.Password,
		"avatar":    b.Avatar,
		"Role":      b.Role,
		"birth_day": b.BirthDay,
		"phone":     b.PhoneNumber,
		"position":  b.Positions,
	})
}

func (a Api) CreateAdmin(c *gin.Context) {
	admin := models.Users{}

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	_, err := time.Parse("2006-01-02", admin.BirthDay)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	if err := a.storage.Admin().CreateUser(context.Background(), admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func (a Api) UpdateAdmin(c *gin.Context) {
	admin := models.Users{}
	if err := c.ShouldBindJSON(admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if err := a.storage.Admin().UpdateUser(context.Background(), admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func (a Api) DeleteAdmin(c *gin.Context) {
	id := c.Param("id")

	err := a.storage.Admin().DeleteUser(c.Request.Context(), id)
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

func (a Api) CreateProject(c *gin.Context) {
	project := models.Project{}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if err := a.storage.Admin().CreateProject(context.Background(), project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (a Api) UpdateProject(c *gin.Context) {
	project := models.Project{}

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	err := a.storage.Admin().UpdateProject(context.Background(), project)
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

func (a Api) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	err := a.storage.Admin().DeleteProject(context.Background(), id)
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

func (a Api) GetUserList(c *gin.Context) {
	res, err := a.storage.Admin().GetUserList(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (a Api) ProjectList(c *gin.Context) {

	res, err := a.storage.Admin().GetProjectList(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (a Api) GetProject(c *gin.Context) {
	id := c.Param("id")
	b, err := a.storage.Admin().GetProject(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
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

//func (a AdminService) CheckTeamLead(c *gin.Context) {
//	check := models.CheckTeamLeadRequest{}
//	if err := c.ShouldBindJSON(check); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"err": errors.New("It isn't Admin's method"),
//		})
//		return
//	}
//
//	t, err := a.storage.Admin().CheckTeamLead(context.Background(), check)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"err": err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, t)
//}

//func (a AdminService) GetUserRole(c *gin.Context) {
//	userId := c.Param("user_id")
//
//	role, err := a.storage.Admin().GetUserRole(context.Background(), userId)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"err": err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, role)
//}

/*
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
	task := models.Task{}

	if err := c.ShouldBindJSON(&task); err != nil {
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

func (a Api) UpdateTask(c *gin.Context) {
	task := models.Task{}
	if err := c.ShouldBindJSON(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	err := a.Repo.UpdateTask(context.Background(), task)
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

func (a Api) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := a.Repo.DeleteTask(context.Background(), id)
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
*/
