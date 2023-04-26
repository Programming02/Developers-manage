package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a Api) ProjectList(c *gin.Context) {

	res, err := a.Repo.ProjectList(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
