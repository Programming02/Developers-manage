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

	r.GET("/task/", c.GetTask)
	r.GET("/commits/", c.GetCommits)

	r.PUT("/up_task/", c.UpdateTask)
	r.PUT("/up_commit/", c.UpdateCommit)

	r.DELETE("/del_task/", c.DeleteTask)
	r.DELETE("/del_commit/", c.DeleteCommit)
}
