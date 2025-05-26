package model

import (
	"gorm.io/gorm"

	"backend/model/work"
)

func InitTable(db *gorm.DB) error {
	// 自动迁移数据库表结构
	err := db.AutoMigrate(
		&work.Work{},
	)
	if err != nil {
		return err
	}
	return nil
}