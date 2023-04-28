package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	middleware "github.com/programming02/osg/middlewere"
)

func InitRouter(a *controller.Api) *gin.Engine {
	r := gin.Default()
	cfg := config.Load()

	r.Use(middleware.Authorizer(cfg))

	r.GET("/admin/:id/", a.AdminService.GetAdmin)
	r.POST("/admin/", a.AdminService.CreateAdmin)
	r.DELETE("/del_admin/:id/", a.AdminService.DeleteAdmin)
	//r.PUT("user_update", a.UpdateAdmin)
	r.GET("/project/:id/", a.AdminService.GetProject)
	r.POST("/project/", a.AdminService.CreateProject)
	r.DELETE("/del_project/:id/", a.AdminService.DeleteProject)
	//r.PUT("up_proj/",)
	//r.POST("/task/", a.AdminService.CreateTask)
	return r
}
