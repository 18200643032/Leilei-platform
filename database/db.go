package database

import (
	"Leilei-platform/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var (
	Read,
	Writer string
	DB *gorm.DB
)

func InitMysql(read, writer string) {
	// 重连保存密码
	Read = read
	Writer = writer
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      read,
		DefaultStringSize:        256,  // string字符默认长度
		DisableDatetimePrecision: true, // 禁止datetime精度 mysql 5.6不支持
		DontSupportRenameIndex:   true, // 重命名索引，把索引删除重建 mysql 5.7不支持
		DontSupportRenameColumn:  true, // 用change重命名列 mysql 8之前支持
	}), &gorm.Config{
		Logger: logger.Default, // 日志等级
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名的命名策略是否是驼峰
		},
	})
	if err != nil {
		return
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(20)                  // 设置连接池
	sqlDb.SetMaxOpenConns(100)                 // 打开连接数
	sqlDb.SetConnMaxIdleTime(time.Second * 30) // 超时时间
	DB = db
	// 主从配置
	_ = DB.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{ // 读操作
			mysql.Open(writer),
		},
		Replicas: []gorm.Dialector{ // 写操作
			mysql.Open(read),
		},
		Policy: dbresolver.RandomPolicy{},
	}))
	err = db.AutoMigrate(&model.User{}, &model.Project{}, &model.Role{}, &model.UserRole{})
	if err != nil {
		fmt.Println("自动创建表结构失败：", err)
		return
	}
}

func NewDBLint() *gorm.DB {
	db := DB
	return db
}
