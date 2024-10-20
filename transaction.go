package db

import "gorm.io/gorm"

// Transaction 事务
//
// tranFunc
//
//	里面只能用参数tx
func Transaction(tranFunc func(tx *gorm.DB) error) error {
	return db.Transaction(tranFunc)
}
