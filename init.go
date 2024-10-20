package db

import (
	"context"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(ctx context.Context, dbPath string) error {
	gormDb, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		AllowGlobalUpdate: false, // 阻止没有条件的更新或删除，会返回ErrMissingWhereClause错误
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		SkipDefaultTransaction: true, // 禁用它，这将获得大约 30%+ 性能提升
		PrepareStmt:            true, // 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
	})
	if err != nil {
		return err
	}
	db = gormDb.WithContext(ctx)
	return nil
}

func GetDb() *gorm.DB {
	return db
}
