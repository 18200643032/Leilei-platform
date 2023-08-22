package dto

type PermissionEnum string

const (
	NORMAL_MENU  PermissionEnum = "NORMAL_MENU"
	USER_MENU    PermissionEnum = "USER_MENU"
	ROLE_MENU    PermissionEnum = "ROLE_MENU"
	PROJECT_MENU PermissionEnum = "PROJECT_MENU"
	SETTING_MENU PermissionEnum = "SETTING_MENU"
	SETTING_OPT  PermissionEnum = "SETTING_OPT"
)

type MenuDTO struct {
	ID    int       `json:"id"`
	Name  string    `json:"name"`
	Icon  string    `json:"icon"`
	Path  *string   `json:"path"`
	Menus []MenuDTO `json:"menus,omitempty"`
}

//menus := map[MenuEnum]Menu{
//COMMON:      {100, "公共组件", "icon-gonggongzujian", "", nil, NORMAL_MENU},
//FILE:        {111, "文件管理", "icon-wenjianguanli", "/common/fileManage", &menus[COMMON], NORMAL_MENU},
//PARAM:       {121, "公共参数", "icon-gonggongcanshu", "/common/commonParam", &menus[COMMON], NORMAL_MENU},
//FUNCTION:    {131, "函数管理", "icon-hanshuguanli", "/common/funcManage", &menus[COMMON], NORMAL_MENU},
//OPERATION:   {141, "操作管理", "icon-caozuoguanli", "/common/operationManage", &menus[COMMON], NORMAL_MENU},
//ENVIRONMENT: {200, "环境中心", "icon-huanjingzhongxin", "", nil, NORMAL_MENU},
//SERVER:      {211, "环境管理", "icon-huanjingguanli", "/envCenter/envManage", &menus[ENVIRONMENT], NORMAL_MENU},
//ENGINE:      {221, "引擎管理", "icon-yinqinguanli", "/envCenter/engineManage", &menus[ENVIRONMENT], NORMAL_MENU},
//DEVICE:      {231, "设备管理", "icon-shebeiguanli", "/envCenter/deviceManage", &menus[ENVIRONMENT], NORMAL_MENU},
//CORE:        {300, "用例中心", "icon-yonglizhongxin", "", nil, NORMAL_MENU},
//API:         {311, "接口管理", "icon-jiekouguanli", "/caseCenter/interfaceManage", &menus[CORE], NORMAL_MENU},
//ELEMENT:     {321, "元素管理", "icon-yuansuguanli", "/caseCenter/elementManage", &menus[CORE], NORMAL_MENU},
//CONTROL:     {326, "控件管理", "icon-kongjianguanli", "/caseCenter/controlManage", &menus[CORE], NORMAL_MENU},
//CASE:        {331, "用例管理", "icon-yongliguanli", "/caseCenter/caseManage", &menus[CORE], NORMAL_MENU},
//TEST:        {400, "计划管理", "icon-jihuaguanli", "", nil, NORMAL_MENU},
//COLLECTION:  {411, "测试集合", "icon-ceshijihe", "/planManage/testCollection", &menus[TEST], NORMAL_MENU},
//PLAN:        {421, "测试计划", "icon-ceshijihua", "/planManage/testPlan", &menus[TEST], NORMAL_MENU},
//RESULT:      {500, "测试追踪", "icon-ceshizhuizong", "", nil, NORMAL_MENU},
//REPORT:      {511, "测试报告", "icon-ceshibaogao", "/report/testReport", &menus[RESULT], NORMAL_MENU},
//SYSTEM:      {600, "系统管理", "icon-xitongguanli", "", nil, NORMAL_MENU},
//USER:        {611, "用户管理", "icon-yonghuguanli", "/systemManage/user", &menus[SYSTEM], USER_MENU},
//ROLE:        {621, "角色管理", "icon-jiaoseguanli", "/systemManage/role", &menus[SYSTEM], ROLE_MENU},
//PROJECT:     {631, "项目管理", "icon-xiangmuguanli", "/systemManage/project", &menus[SYSTEM], PROJECT_MENU},
//}
