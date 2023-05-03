package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	middleware "github.com/programming02/osg/middlewere"
)

func AdminRouter(a controller.Api) {
	r := gin.Default()
	cfg, err := config.Load()
	if err != nil {
		return
	}

	r.Use(middleware.Authorizer(cfg))

	r.GET("/get_admin/:id/", a.CreateAdmin)
	r.GET("/get_project/", a.GetProject)
	r.GET("/get_users/", a.GetUserList)
	r.GET("/get_projects/", a.ProjectList)

	r.POST("/admin/", a.CreateAdmin)
	r.POST("/project/", a.CreateProject)

	r.PUT("/up_user/", a.UpdateAdmin)
	r.PUT("/up_project/", a.UpdateProject)

	r.DELETE("/del_admin/:id/", a.DeleteAdmin)
	r.DELETE("/del_project/", a.DeleteProject)

}
