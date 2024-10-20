package db

import "gorm.io/gorm"

// Hooks
//
// 事务完成之前该事务中所作的更改是不可见的，如果您的钩子返回了任何错误，则修改将被回滚

//                      _
//   ___ _ __ ___  __ _| |_ ___
//  / __| '__/ _ \/ _` | __/ _ \
// | (__| | |  __/ (_| | ||  __/
//  \___|_|  \___|\__,_|\__\___|

// 开始事务

func (d *DbBase) BeforeSave(tx *gorm.DB) (err error) {
	return
}

func (d *DbBase) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

// 关联前的 save

// 插入记录至 db

// 关联后的 save

func (d *DbBase) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func (d *DbBase) AfterSave(tx *gorm.DB) (err error) {
	return
}

// 提交或回滚事务

//                  _       _
//  _   _ _ __   __| | __ _| |_ ___
// | | | | '_ \ / _` |/ _` | __/ _ \
// | |_| | |_) | (_| | (_| | ||  __/
//  \__,_| .__/ \__,_|\__,_|\__\___|
//       |_|

// 开始事务

// BeforeSave

func (d *DbBase) BeforeUpdate(tx *gorm.DB) (err error) {
	return
}

// 关联前的 save

// 更新 db

// 关联后的 save

func (d *DbBase) AfterUpdate(tx *gorm.DB) (err error) {
	return
}

// AfterSave

// 提交或回滚事务

//      _      _      _
//   __| | ___| | ___| |_ ___
//  / _` |/ _ \ |/ _ \ __/ _ \
// | (_| |  __/ |  __/ ||  __/
//  \__,_|\___|_|\___|\__\___|

// 开始事务

func (d *DbBase) BeforeDelete(tx *gorm.DB) (err error) {
	return
}

// 删除 db 中的数据

func (d *DbBase) AfterDelete(tx *gorm.DB) (err error) {
	return
}

// 提交或回滚事务

//   __ _           _
//  / _(_)_ __   __| |
// | |_| | '_ \ / _` |
// |  _| | | | | (_| |
// |_| |_|_| |_|\__,_|

// 从 db 中加载数据

// Preloading (eager loading)

func (d *DbBase) AfterFind(tx *gorm.DB) (err error) {
	return
}
