package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// 分页查询
func PaginateAndSearch(db *gorm.DB, model interface{}, keyword string, page, pageSize int) (int, error) {
	var total int64
	offset := (page - 1) * pageSize
	err := db.Model(model).Where("name LIKE ?", "%"+keyword+"%").Count(&total).Error
	if err != nil {
		return 0, err
	}

	err = db.Where("name LIKE ?", "%"+keyword+"%").Offset(offset).Limit(pageSize).Find(model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}

	return int(total), nil
}

func Page(c *gin.Context) (string, int, int) {
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
	return name, pageInt, pageSizeInt

}
