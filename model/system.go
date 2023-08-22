package model

import (
	"Leilei-platform/global"
	"github.com/gofrs/uuid"
)

type User struct {
	global.INIT_MODEL
	Username    string `json:"userName" gorm:"default:系统管理员;comment:用户名"` // 用户登录名
	Account     string `json:"account" gorm:"index;comment:用户名登录名称"`
	Password    string `json:"-"  gorm:"comment:用户登录密码"`            // 用户登录密码
	LastProject string `json:"lastProject" gorm:"comment:上一次访问的项目"` // 用户昵称
	Mobile      string `json:"mobile"  gorm:"comment:用户手机号"`        // 用户手机号
	Email       string `json:"email"  gorm:"comment:用户邮箱"`          // 用户邮箱
	Status      string `json:"status" gorm:"comment:状态"`
}

func (User) TableName() string {
	return "user"
}

type Project struct {
	global.INIT_MODEL
	Name         string    `json:"name" gorm:"comment:项目名"`
	Description  string    `json:"description" gorm:"comment:项目描述"`
	ProjectAdmin uuid.UUID `json:"project_admin" gorm:"comment:项目管理员"`
	Status       string    `json:"status" gorm:"default:启用;comment:状态"`
}

func (Project) TableName() string {
	return "project"
}

type Role struct {
	global.INIT_MODEL
	Name      string    `json:"name" gorm:"comment:角色名"`
	ProjectId uuid.UUID `json:"projectId" gorm:"comment:项目id"`
}

func (Role) TableName() string {
	return "role"
}

type Permission struct {
	ID   string `json:"id"` //
	Name string `json:"name" gorm:"comment:权限名称"`
}

func (Permission) TableName() string {
	return "permission"
}

type UserRole struct {
	global.INIT_MODEL
	UserId string    `json:"user_id" gorm:"comment:用户id"`
	RoleId uuid.UUID `json:"role_id" gorm:"comment:角色id"`
}

func (UserRole) TableName() string {
	return "user_role"
}

type UserProject struct {
	global.INIT_MODEL
	UserId    string `json:"user_id" gorm:"comment:用户id"`
	ProjectId string `json:"project_id" gorm:"comment:项目id"`
}

func (UserProject) TableName() string {
	return "user_project"
}

type RolePermission struct {
	global.INIT_MODEL
	RoleId       string `json:"role_id" gorm:"comment:角色id"`
	PermissionId string `json:"permission_id" gorm:"comment:权限id"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}

// 注册
type Register struct {
	Username string `json:"userName" example:"用户名"`
	Password string `json:"passWord" example:"密码"`
	Account  string `json:"account" example:"用户名"`
	Mobile   string `json:"phone" example:"电话号码"`
	Email    string `json:"email" example:"电子邮箱"`
}

// 登录
type Login struct {
	Account  string `json:"account"`  // 用户名
	Password string `json:"password"` // 密码
}

// 添加项目
type AddProject struct {
	Description  string `json:"description"`
	Name         string `json:"name"`
	ProjectAdmin string `json:"projectAdmin"`
}
