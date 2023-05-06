package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	middleware "github.com/programming02/osg/middlewere"
)

func AdminRouter(a *controller.Api, r gin.IRouter) {
	cfg, err := config.Load()
	if err != nil {
		return
	}

	r.Use(middleware.Authorizer(cfg))

	r.GET("/get-admin/:id/", a.GetAdmin)
	r.GET("/get-project/", a.GetProject)
	r.GET("/get-users/", a.GetUserList)
	r.GET("/get-projects/", a.ProjectList)

	r.POST("/admin/", a.CreateAdmin)
	r.POST("/project/", a.CreateProject)

	r.PUT("/up-user/", a.UpdateAdmin)
	r.PUT("/up-project/", a.UpdateProject)

	r.DELETE("/del-admin/:id/", a.DeleteAdmin)
	r.DELETE("/del-project/", a.DeleteProject)
}
