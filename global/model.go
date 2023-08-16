package global

import (
	"time"
)

type INIT_MODEL struct {
	ID         int64     `json:"id" gorm:"primarykey"` // 主键ID
	CreateTime time.Time `json:"createTime"`           // 创建时间
	UpdateTime time.Time `json:"updateTime"`           // 更新时间
}
