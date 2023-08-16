package system

import (
	"github.com/gin-gonic/gin"
	"Leilei-platform/utils"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"Leilei-platform/database"
	"Leilei-platform/middleware"
	"context"
	"Leilei-platform/model"
	"Leilei-platform/dto"
	"fmt"
	"time"
	"Leilei-platform/global"
	"Leilei-platform/public/response"
	"Leilei-platform/public/e"
)

//登录

func Login(c *gin.Context) {
	var loginInfo model.Login

	appG := response.Gin{C: c}
	err := c.BindJSON(&loginInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}
	// 查询用户信息
	var user model.User
	if err = database.DB.Where("account = ?", loginInfo.Account).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
		return
	}
	// 验证密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "密码错误"})
		return
	}

	// 生成JWT Token
	token, err := utils.GenerateToken(user.ID, user.Account, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token错误"})
		return
	}
	fmt.Println(token)

	// 更新Token字段并返回成功提示信息
	rdb := middleware.NewRedisClient()
	ctx := context.Background()
	err = rdb.Set(ctx, "token", token, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "服务器错误"})
		return
	}
	// 将user数据复制到loginres中
	loginres := dto.UserDTO{
		User:  user,
		Token: token,
		// 其他字段赋值...
	}
	fmt.Println(loginres)
	appG.Result(http.StatusOK, e.SUCCESS, loginres)
	//c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": token, "data": user})
}

//注册
func Register(c *gin.Context) {
	var user model.Register
	appG := response.Gin{C: c}
	err := c.BindJSON(&user)
	if err != nil {
		appG.Result(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	var count int64
	database.DB.Model(&model.User{}).Where("account = ?", user.Account).Count(&count)
	if count > 0 {
		appG.Result(http.StatusBadRequest, e.ERROR_EXIST_TAG, nil)
		return
	}
	//密码加密
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码服务器错误"})
		return
	}
	// 创建用户对象并赋值
	newUser := &model.User{
		Account:    user.Account,
		Password:   string(passwordHash),
		Email:      user.Email,
		INIT_MODEL: global.INIT_MODEL{CreateTime: time.Now(), UpdateTime: time.Now()}, // 设置为当前时间

		// 其他字段赋值...
	}
	// 存储用户信息
	user.Password = string(passwordHash)
	if err = database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	appG.Result(http.StatusOK, e.SUCCESS, nil)

}
func UserAll(c *gin.Context) {
	var user []model.User
	appG := response.Gin{C: c}
	err := c.Bind(&user)
	if err != nil {
		appG.Result(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	if err = database.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询错误"})
		return
	}
	appG.Result(http.StatusOK, e.SUCCESS, user)
}

func GetUserInfo(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	appG := response.Gin{C: c}
	err := c.Bind(&user)
	if err != nil {
		appG.Result(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	if err = database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询错误"})
		return
	}
	loginres := dto.UserDTO{
		User: user,
		// 其他字段赋值...
	}
	appG.Result(http.StatusOK, e.SUCCESS, loginres)
}

func SwitchProject(c *gin.Context) {
	// 获取项目ID
	appG := response.Gin{C: c}
	requestData := dto.SwitchProjectDTO{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if requestData.LastProject == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lastProject不能为空"})
		return
	}

	// 根据项目ID更新
	err := database.DB.Model(&model.User{}).Where("id = ?", requestData.ID).Update("last_project", requestData.LastProject).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 查询更新后的用户信息
	var updatedUser model.User
	err = database.DB.First(&updatedUser, requestData.ID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询错误"})
		return
	}

	appG.Result(http.StatusOK, e.SUCCESS, updatedUser)
}

//更新密码
func UpdatePassword(c *gin.Context) {
	// 获取项目ID
	requestData := dto.UpdatePasswordDTO{}
	err := c.BindJSON(&requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}
	appG := response.Gin{C: c}
	var user model.User

	infoerr := database.DB.Where("id = ?", requestData.UserId).First(&user).Error
	if infoerr != nil {
		c.JSON(http.StatusBadRequest, infoerr.Error())
		return
	}
	oldErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestData.OldPassword))
	if oldErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "原密码错误"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(requestData.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码服务器错误"})
		return
	}

	// 根据项目ID更新
	updateErr := database.DB.Model(&model.User{}).Where("id = ?", requestData.UserId).Update("password", passwordHash).Error
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appG.Result(http.StatusOK, e.SUCCESS, nil)
}

//更新资料
func UpdateInfo(c *gin.Context) {
	// 获取项目ID
	appG := response.Gin{C: c}
	requestData := dto.UserDTO{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user model.User
	infoerr := database.DB.Where("id = ?", requestData.ID).First(&user).Error
	if infoerr != nil {
		c.JSON(http.StatusBadRequest, infoerr.Error())
		return
	}

	// 根据项目ID更新
	if requestData.Username != "" {
		user.Username = requestData.Username
	}
	if requestData.Email != "" {
		user.Email = requestData.Email
	}
	user.UpdateTime = time.Now()
	err := database.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	appG.Result(http.StatusOK, e.SUCCESS, nil)
}

//获取项目下的所有角色
func GetProjectRole(c *gin.Context) {
	var roles []model.Role
	appG := response.Gin{C: c}
	projectID := c.Query("projectId")
	//err := c.Bind(&roles)
	if err := database.DB.Where("project_id = ?", projectID).Find(&roles).Error; err != nil {
		appG.Result(http.StatusBadRequest, e.ERROR, nil)
		return
	}
	appG.Result(http.StatusOK, e.SUCCESS, roles)
}
