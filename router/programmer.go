package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	middleware "github.com/programming02/osg/middlewere"
)

func ProgrammerRouter(c *controller.Api, r gin.IRouter) {
	cfg, err := config.Load()
	if err != nil {
		return
	}

	r.Use(middleware.Authorizer(cfg))

	r.POST("/task/", c.CreateTask)
	r.POST("/commit/", c.CreateCommit)
	r.POST("/attendance/", c.CreateAttendance)

	r.GET("/get-task/", c.GetTask)
	r.GET("/get-commits/", c.GetCommit)

	r.PUT("/up-task/", c.UpdateTask)
	r.PUT("/up-commit/", c.UpdateCommit)

	r.DELETE("/del-task/", c.DeleteTask)
	r.DELETE("/del-commit/", c.DeleteCommit)
}
