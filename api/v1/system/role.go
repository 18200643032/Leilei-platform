package system

import (
	"Leilei-platform/database"
	"Leilei-platform/dto"
	"Leilei-platform/global"
	"Leilei-platform/model"
	"Leilei-platform/public/e"
	"Leilei-platform/public/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoleAll(c *gin.Context) {
	var roles []model.Role
	appG := response.Gin{C: c}
	//projectID := c.Query("projectId")
	//err := c.Bind(&roles)
	name, pageInt, pageSizeInt := global.Page(c)
	total, err := global.PaginateAndSearch(database.DB, &roles, name, pageInt, pageSizeInt)
	if err != nil {
		appG.Result(http.StatusBadRequest, e.ERROR, nil)
	}
	roleResponse := dto.RoleDTO{
		Total: int64(total),
		List:  roles,
	}
	if err := database.DB.Find(&roles).Error; err != nil {
		appG.Result(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Result(http.StatusOK, e.SUCCESS, roleResponse)
}

func GetRoleUser(c *gin.Context) {
	var users []model.User
	appG := response.Gin{C: c}
	roleId := c.Query("roleId")
	//err := c.Bind(&roles)
	err := database.DB.Preload("UserRole").Where("user_role.role_id = ?", roleId).Find(&users).Error
	name, pageInt, pageSizeInt := global.Page(c)
	total, err := global.PaginateAndSearch(database.DB, &users, name, pageInt, pageSizeInt)
	if err != nil {
		appG.Result(http.StatusBadRequest, e.ERROR, nil)
	}
	roleResponse := dto.RoleUser{
		Total: int64(total),
		List:  users,
	}
	if err := database.DB.Find(&users).Error; err != nil {
		appG.Result(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Result(http.StatusOK, e.SUCCESS, roleResponse)
}
