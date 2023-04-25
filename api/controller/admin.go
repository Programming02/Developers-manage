package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/moduls"
	// "go get -u github.com/swaggo/files"
	// "go get -u github.com/swaggo/gin-swagger"
	"net/http"
)

func (a Api) GetAdmin(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
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
