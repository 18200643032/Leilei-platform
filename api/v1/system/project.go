package system

import (
	"github.com/gin-gonic/gin"
	"Leilei-platform/model"
	"time"
	"Leilei-platform/database"
	"net/http"
	"Leilei-platform/global"
	"github.com/gofrs/uuid"
	"strconv"
	"fmt"
	"Leilei-platform/dto"
	"Leilei-platform/public/response"
	"Leilei-platform/public/e"
)

func AddProject(c *gin.Context) {
	var project model.Project
	appG := response.Gin{C: c}
	err := c.BindJSON(&project)
	if err != nil {
		appG.Result(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	var count int64
	database.DB.Model(&model.Project{}).Where("name = ?", project.Name).Count(&count)
	if count > 0 {
		appG.Result(http.StatusBadRequest, e.ERROR_EXIST_TAG, nil)
		return
	}
	newUUID, _ := uuid.NewV4()
	// 创建用户对象并赋值
	newProject := &model.Project{
		Name:         project.Name,
		Description:  project.Description,
		ProjectAdmin: newUUID,
		INIT_MODEL:   global.INIT_MODEL{CreateTime: time.Now(), UpdateTime: time.Now()}, // 设置为当前时间

		// 其他字段赋值...
	}
	// 存储用户信息
	if err = database.DB.Create(&newProject).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	appG.Result(http.StatusOK, e.SUCCESS, nil)
}

//获取项目详情
func GetProject(c *gin.Context) {
	var project []model.Project
	appG := response.Gin{C: c}

	projectID := c.Query("projectId")
	//err := c.Bind(&roles)
	if err := database.DB.Where("project_admin = ?", projectID).First(&project).Error; err != nil {
		appG.Result(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Result(http.StatusOK, e.SUCCESS, project)
}

//获取所有项目
func GetAllProjects(c *gin.Context) {

	appG := response.Gin{C: c}
	projectID := c.Query("projectId")
	if projectID == "" {
		var projects []model.Project

		// 获取分页参数
		page := c.Query("page")
		pageSize := c.Query("pageSize")
		name := c.Query("name")
		// 设置默认值
		if name == "" {
			name = ""
		}
		if page == "" {
			page = "1"
		}
		if pageSize == "" {
			pageSize = "10"
		} // 将字符串类型的分页参数转换为整数类型
		pageInt, _ := strconv.Atoi(page)
		pageSizeInt, _ := strconv.Atoi(pageSize)

		// 构建查询对象
		total, err := global.PaginateAndSearch(database.DB, &projects, name, pageInt, pageSizeInt)
		if err != nil {
			appG.Result(http.StatusBadRequest, e.ERROR, nil)
		}
		for _, project := range projects {
			fmt.Println(project)
		}
		projectResponse := dto.ProjectResponse{
			Total: int64(total),
			List:  projects,
		}

		appG.Result(http.StatusOK, e.SUCCESS, projectResponse)
	} else {
		var projects model.Project
		if err := database.DB.Where("id = ?", projectID).First(&projects).Error; err != nil {
			appG.Result(http.StatusBadRequest, e.ERROR, nil)
			return
		}
		appG.Result(http.StatusOK, e.SUCCESS, projects)
	}
}

//删除项目
func DeleteProject(c *gin.Context) {
	// 获取项目ID
	appG := response.Gin{C: c}
	projectID := c.Query("id")
	var count int64
	database.DB.Model(&model.Project{}).Where("id = ?", projectID).Count(&count)
	if count < 1 {
		appG.Result(http.StatusBadRequest, e.ERROR_NOT_EXIST_TAG, nil)
	}

	// 根据项目ID删除项目信息
	if err := database.DB.Model(&model.Project{}).Where("id = ?", projectID).Update("status", "停用").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appG.Result(http.StatusOK, e.SUCCESS, nil)
}

//恢复项目
func RecoverProject(c *gin.Context) {
	// 获取项目ID
	appG := response.Gin{C: c}
	projectID := c.Query("id")
	// 根据项目ID恢复项目信息
	if err := database.DB.Model(&model.Project{}).Where("id = ?", projectID).Update("status", "启用").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appG.Result(http.StatusOK, e.SUCCESS, nil)
}
