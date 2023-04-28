package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	middleware "github.com/programming02/osg/middlewere"
)

func RegisterROuter(c controller.Api) {
	r := gin.Default()
	cfg := config.Load()

	r.Use(middleware.Authorizer(cfg))

	r.POST("/login/", c.Login)
}
