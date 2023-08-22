package system

import (
	"Leilei-platform/dto"
	"Leilei-platform/public/e"
	"Leilei-platform/public/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取项目详情
func GetMenu(c *gin.Context) {
	fileManagePath := "/common/fileManage"
	commonParamPath := "/common/commonParam"
	funcManagePath := "/common/funcManage"
	operationManagePath := "/common/operationManage"
	envManagePath := "/envCenter/envManage"
	engineManagePath := "/envCenter/engineManage"
	deviceManagePath := "/envCenter/deviceManage"
	interfaceManagePath := "/caseCenter/interfaceManage"
	elementManagePath := "/caseCenter/elementManage"
	controlManagePath := "/caseCenter/controlManage"
	caseManagePath := "/caseCenter/caseManage"
	testCollectionPath := "/planManage/testCollection"
	testPlanPath := "/planManage/testPlan"
	testReportPath := "/report/testReport"
	userPath := "/systemManage/user"
	rolePath := "/systemManage/role"
	projectPath := "/systemManage/project"

	menu := []dto.MenuDTO{
		{
			ID:   100,
			Name: "公共组件",
			Icon: "icon-gonggongzujian",
			Path: nil,
			Menus: []dto.MenuDTO{
				{
					ID:    111,
					Name:  "文件管理",
					Icon:  "icon-wenjianguanli",
					Path:  &fileManagePath,
					Menus: nil,
				},
				{
					ID:    121,
					Name:  "公共参数",
					Icon:  "icon-gonggongcanshu",
					Path:  &commonParamPath,
					Menus: nil,
				},
				{
					ID:    131,
					Name:  "函数管理",
					Icon:  "icon-hanshuguanli",
					Path:  &funcManagePath,
					Menus: nil,
				},
				{
					ID:    141,
					Name:  "操作管理",
					Icon:  "icon-caozuoguanli",
					Path:  &operationManagePath,
					Menus: nil,
				},
			},
		},
		{
			ID:   200,
			Name: "环境中心",
			Icon: "icon-huanjingzhongxin",
			Path: nil,
			Menus: []dto.MenuDTO{
				{
					ID:    211,
					Name:  "环境管理",
					Icon:  "icon-huanjingguanli",
					Path:  &envManagePath,
					Menus: nil,
				},
				{
					ID:    221,
					Name:  "引擎管理",
					Icon:  "icon-yinqinguanli",
					Path:  &engineManagePath,
					Menus: nil,
				},
				{
					ID:    231,
					Name:  "设备管理",
					Icon:  "icon-shebeiguanli",
					Path:  &deviceManagePath,
					Menus: nil,
				},
			},
		},
		{
			ID:   300,
			Name: "用例中心",
			Icon: "icon-yonglizhongxin",
			Path: nil,
			Menus: []dto.MenuDTO{
				{
					ID:    311,
					Name:  "接口管理",
					Icon:  "icon-jiekouguanli",
					Path:  &interfaceManagePath,
					Menus: nil,
				},
				{
					ID:    321,
					Name:  "元素管理",
					Icon:  "icon-yuansuguanli",
					Path:  &elementManagePath,
					Menus: nil,
				},
				{
					ID:    326,
					Name:  "控件管理",
					Icon:  "icon-kongjianguanli",
					Path:  &controlManagePath,
					Menus: nil,
				},
				{
					ID:    331,
					Name:  "用例管理",
					Icon:  "icon-yongliguanli",
					Path:  &caseManagePath,
					Menus: nil,
				},
			},
		},
		{
			ID:   400,
			Name: "计划管理",
			Icon: "icon-jihuaguanli",
			Path: nil,
			Menus: []dto.MenuDTO{
				{
					ID:    411,
					Name:  "测试集合",
					Icon:  "icon-ceshijihe",
					Path:  &testCollectionPath,
					Menus: nil,
				},
				{
					ID:    421,
					Name:  "测试计划",
					Icon:  "icon-ceshijihua",
					Path:  &testPlanPath,
					Menus: nil,
				},
			},
		},
		{
			ID:   500,
			Name: "测试追踪",
			Icon: "icon-ceshizhuizong",
			Path: nil,
			Menus: []dto.MenuDTO{
				{
					ID:    511,
					Name:  "测试报告",
					Icon:  "icon-ceshibaogao",
					Path:  &testReportPath,
					Menus: nil,
				},
			},
		},
		{
			ID:   600,
			Name: "系统管理",
			Icon: "icon-xitongguanli",
			Path: nil,
			Menus: []dto.MenuDTO{
				{
					ID:    611,
					Name:  "用户管理",
					Icon:  "icon-yonghuguanli",
					Path:  &userPath,
					Menus: nil,
				},
				{
					ID:    621,
					Name:  "角色管理",
					Icon:  "icon-jiaoseguanli",
					Path:  &rolePath,
					Menus: nil,
				},
				{
					ID:    631,
					Name:  "项目管理",
					Icon:  "icon-xiangmuguanli",
					Path:  &projectPath,
					Menus: nil,
				},
			},
		},
	}

	appG := response.Gin{C: c}

	//projectID := c.Query("projectId")
	//err := c.Bind(&roles)
	appG.Result(http.StatusOK, e.SUCCESS, menu)
}
