package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open("root:imh@2025@tcp(49.235.100.155)/db1?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束生成
	})
	if err != nil {
		panic(err)
	}
	DB = db
}

func GetDb() *gorm.DB {
	return DB
}
