package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	middleware "github.com/programming02/osg/middlewere"
)

func AdminRouter(a controller.Api) {
	r := gin.Default()
	cfg := config.Load()

	r.Use(middleware.Authorizer(cfg))

	r.GET("/get_admin/:id/", a.AdminService.CreateAdmin)
	r.GET("/get_project/", a.AdminService.GetProject)
	r.GET("/get_users/", a.AdminService.GetUserList)
	r.GET("/get_projects/", a.AdminService.ProjectList)

	r.POST("/admin/", a.AdminService.CreateAdmin)
	r.POST("/project/", a.AdminService.CreateProject)

	r.PUT("/up_user/", a.AdminService.UpdateAdmin)
	r.PUT("/up_project/", a.AdminService.UpdateProject)

	r.DELETE("/del_admin/:id/", a.AdminService.DeleteAdmin)
	r.DELETE("/del_project/", a.AdminService.DeleteProject)

}
