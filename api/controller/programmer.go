package controller

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/programming02/osg/api/models"
	"github.com/programming02/osg/jwt"
	"net/http"
)

func (a Api) CreateTask(c *gin.Context) {
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

	role, err := a.storage.Programmer().UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: task.ProjectId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	if role != "team_lead" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "You're not team_lead"})
	}

	if err := a.storage.Programmer().CreateTask(context.Background(), models.CreateTaskRequestModel{}); err != nil {
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

	role, err := a.storage.Programmer().UserRole(context.Background(), models.UserRole{
		UserId: user.UserID.String(),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	if role != "team_lead" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "You're not team_lead"})
	}

	err = a.storage.Programmer().UpdateTask(context.Background(), task)
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

func (a Api) DeleteTask(c *gin.Context) {
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

	task, err := a.storage.Programmer().GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	role, err := a.storage.Programmer().UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: task.ProjectId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	if role != "team_lead" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "You're not team_lead"})
	}

	err = a.storage.Programmer().DeleteTask(context.Background(), id)
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

func (a Api) GetTask(c *gin.Context) {
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

	t, err := a.storage.Programmer().GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	_, err = a.storage.Programmer().UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: t.ProjectId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"err": errors.New("you don't have access")})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	task, err := a.storage.Programmer().GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"ok": task})
}

func (a Api) CreateCommit(c *gin.Context) {
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

	t, err := a.storage.Programmer().GetTask(context.Background(), commit.TaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	_, err = a.storage.Programmer().UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: t.ProjectId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"err": errors.New("you don't have access")})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"err": err.Error(),
	})

	err = a.storage.Programmer().CreateCommit(context.Background(), commit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func (a Api) UpdateCommit(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	com := models.Commit{}
	if err := c.ShouldBindJSON(com); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	err = a.storage.Programmer().UpdateCommit(context.Background(), com, user.UserID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (a Api) GetCommitList(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	taskID := c.Param("task_id")
	if err := c.ShouldBindJSON(taskID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": errors.New("there's mistake on 'id'")})
	}

	t, err := a.storage.Programmer().GetTask(context.Background(), taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	_, err = a.storage.Programmer().UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: t.ProjectId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"err": errors.New("you don't have access")})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"err": err.Error(),
	})

	res, err := a.storage.Programmer().GetCommitList(context.Background(), taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, res)
}

func (a Api) DeleteCommit(c *gin.Context) {
	id := c.Param("id")

	if err := c.ShouldBindJSON(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "can't bind id" + err.Error()})
	}

	if err := a.storage.Programmer().DeleteCommit(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "can't delete comment" + err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})

	/* user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	id := c.Param("id")
	if _, err = uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": errors.New("there's a mistake on id "),
		})
	}

	t, err := a.storage.Programmer().GetTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	_, err = a.storage.Programmer().UserRole(context.Background(), models.UserRole{
		UserId:    user.UserID.String(),
		ProjectId: t.ProjectId,
	})
	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, gin.H{"err": errors.New("you don't have access")})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"err": err.Error(),
	})

	err = a.storage.Programmer().DeleteCommit(context.Background(), )
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})  */
}

func (a Api) CreateAttendance(c *gin.Context) {
	user, err := jwt.ExtractTokenMetadata(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	types := c.Param("type")
	if types != "came" && types != "gone" {
		c.JSON(http.StatusBadRequest, gin.H{"err": errors.New("err with type")})
	}

	err = a.storage.Programmer().CreateAttendance(context.Background(), models.Attendance{
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

func (a Api) GetAttendanceList(c *gin.Context) {

}

func (a Api) UserRole(ctx context.Context, role models.UserRole) (string, error) {
	return a.storage.Programmer().UserRole(ctx, role)
}
