package router

import (
	"github.com/gin-gonic/gin"
	"github.com/programming02/osg/api/controller"
)

func InitRouter(a *controller.Api) *gin.Engine {
	r := gin.Default()

	r.GET("/admin/id", a.GetAdmin)
	r.POST("/admin", a.CreateAdmin)
	r.DELETE("/del_admin", a.DeleteAdmin)
	//r.POST("user_update", a.UpdateAdmin)
	r.POST("/project", a.CreateProject)
	r.DELETE("/del_project", a.DeleteProject)
	r.POST("/task", a.CreateTask)
	return r
}
