package router

import (
	"Leilei-platform/middleware"
	"net/http"
	"github.com/gin-gonic/gin"
	"Leilei-platform/api/v1/system"
)

func SystemRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]any{"code": "200", "msg": "ok"})
	})
	v := r.Group("/autotest")
	{
		v.POST("/register", system.Register)
		v.POST("/login", system.Login)
	}
	vv := r.Group("/autotest/user")
	{
		vv.GET("/info/:id", system.GetUserInfo)
		vv.POST("/switch/project", system.SwitchProject)
		vv.POST("/update/password", system.UpdatePassword)
		vv.POST("/update/info", system.UpdateInfo)
		vv.GET("/all", system.UserAll)

	}
	//r.POST("/login", system.Login)
	v2 := r.Group("/api/v1")
	v2.Use(middleware.AuthMiddleware())
	{
		v2.POST("/project/add", system.AddProject)

		v2.DELETE("/project", system.DeleteProject)
		v2.GET("/project", system.GetAllProjects)
		v2.POST("/project/recover", system.RecoverProject)

		v2.GET("/project/role/list", system.GetProjectRole)

	}

	return r

}
