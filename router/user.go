package router

import (
	"Leilei-platform/api/v1/system"
	"Leilei-platform/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
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
	systemRouter := r.Group("/autotest/user")
	systemRouter.Use(middleware.AuthMiddleware())
	{
		systemRouter.GET("/info/:id", system.GetUserInfo)
		systemRouter.POST("/switch/project", system.SwitchProject)
		systemRouter.POST("/update/password", system.UpdatePassword)
		systemRouter.POST("/update/info", system.UpdateInfo)
		systemRouter.GET("/all", system.UserAll)

	}
	vvv := r.Group("autotest/role")
	{
		vvv.GET("/list", system.GetRoleAll)
		vvv.GET("/user/list", system.GetRoleUser)
	}
	vvvv := r.Group("autotest")
	{
		vvvv.GET("menu/list", system.GetMenu)
		vvvv.GET("/project/user/:id", system.ProjectUser)
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
