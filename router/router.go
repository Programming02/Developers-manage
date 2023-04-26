package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
)

func InitRouter(a *controller.Api) *gin.Engine {
	r := gin.Default()

	r.GET("/admin/:id/", a.GetAdmin)
	r.POST("/admin/", a.CreateAdmin)
	r.DELETE("/del_admin/:id/", a.DeleteAdmin)
	//r.PUT("user_update", a.UpdateAdmin)
	r.GET("/project/:id/", a.GetProject)
	r.POST("/project/", a.CreateProject)
	r.DELETE("/del_project/:id/", a.DeleteProject)
	//r.PUT("up_proj/",)
	r.POST("/task/", a.CreateTask)
	return r
}
