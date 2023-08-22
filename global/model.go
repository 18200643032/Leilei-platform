package global

type INIT_MODEL struct {
	ID         int64 `json:"id" gorm:"primarykey"` // 主键ID
	CreateTime int64 `json:"createTime"`           // 创建时间
	UpdateTime int64 `json:"updateTime"`           // 更新时间
}
