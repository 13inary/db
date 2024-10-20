package db

import (
	"time"
)

// dbBase 表名为db_bases
//
// 说明：
// 所有标识符转化为下划线形式
// 框架内置模型：gorm.Model
//
// 用法：
// 1. 嵌套该结构体
// 2. 主键字段们添加标签
type DbBase struct {
	// ID默认为表主键
	ID uint `gorm:"primaryKey"`

	// 被自动管理字段
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除，将记录标记为已删除
}
