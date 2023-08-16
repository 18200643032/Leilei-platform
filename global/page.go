package global

import (
	"gorm.io/gorm"
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
