package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	middleware "github.com/programming02/osg/middlewere"
)

func ProgrammerRouter(c *controller.Api) {
	r := gin.Default()
	cfg := config.Load()

	r.Use(middleware.Authorizer(cfg))

	r.GET("/task/", c.ProgrammerService.GetTask)
	r.GET("/commits/", c.ProgrammerService.GetCommits)

	r.PUT("/up_task/", c.ProgrammerService.UpdateTask)
	r.PUT("/up_commit/", c.ProgrammerService.UpdateCommit)

	r.DELETE("/del_task/", c.ProgrammerService.DeleteTask)
	r.DELETE("/del_commit/", c.ProgrammerService.DeleteCommit)
}
