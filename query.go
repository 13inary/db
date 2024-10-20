package main

import (
	"errors"

	"gorm.io/gorm"
)

// Model 指定使用那个表。这里指定后，后面的Find之类可以使用其他结构体
//
// table
//
//	&user{}
func Model(table any) *gorm.DB {
	return db.Model(table)
}

func Table() {
}

// Take 随机获取表的一条记录，返回：是否找到、错误
//
// table 空结构体&user{}
func Take(table any) (bool, error) {
	result := db.Take(table)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

// First id排序后获取第一条记录，返回：是否找到、错误
//
// table 空结构体&user{}
func First(table any) (bool, error) {
	result := db.First(table)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

// Last id排序后获取最后一条记录，返回：是否找到、错误
//
// table 空结构体&user{}
func Last(table any) (bool, error) {
	result := db.Last(table)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

// Find 返回单条记录或多条记录。找不到不会返回错误
//
// tables 空切片&[]user{}
func Find(tables any) error {
	result := db.Find(tables)
	return result.Error
}

// Rows 逐行读取数据到内存
func Rows() {
}

// Scan 和Find类似。但是还可以结合Select将数据保存到非数据库模型，且可以选择部分字段
//
// anyModel 空切片&[]user{}。任何模型都可以
func Scan(anyModel any) error {
	result := db.Scan(anyModel)
	return result.Error
}

// FirstOrInit
//
//	若找到记录：返回记录
//	若找不到记录：返回一个初始化记录
func FirstOrInit() {
}

// Attrs 结合FirstOrInit等函数，不影响查询SQL
//
//	若找到记录：忽略这个
//	若没找到记录：使用这个
func Attrs() {
}

// Assign 结合FirstOrInit等函数，不影响查询SQL
//
//	不管有没有找到纪律：都使用这个
func Assign() {
}

// FirstOrCreate
//
//	若找到记录：返回记录
//	若找不到记录：创建纪律、返回记录
func FirstOrCreate() {
}

// FindInBatches 分批次查询并处理记录
func FindInBatches(tables any, batchSize int, fc func(tx *gorm.DB, batch int) error) error {
	retult := db.FindInBatches(tables, batchSize, fc)
	return retult.Error
	// fc func(tx *gorm.DB, batch int) error 内容：
	//for _, result := range tables {
	//	// 对批中的每条记录进行操作
	//}

	//// 保存对当前批记录的修改
	//tx.Save(&tables)

	//// tx.RowsAffected 提供当前批处理中记录的计数（the count of records in the current batch）
	//// 'batch' 变量表示当前批号（the current batch number）

	//// 返回 error 将阻止更多的批处理
	//return nil
}

// Pluck 返回多个记录的单个列
func Pluck(column string, table any) *gorm.DB {
	return db.Pluck(column, table)
}

// Scopes 复用多个条件判断
//
// whereFunc
//
//	注意这里的whereFunc中是传入*gorm.DB而不是一个*gorm.DB的实例
func Scopes(whereFunc ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	return db.Scopes(whereFunc...)
}

// Where
//
// condition 和 args
//
//	"name = ?"                     ""
//	"name <> ?"                    ""
//	"name IN ?"                    []string{""}
//	"name LIKE ?"                  "%jin%"
//	"name = ? AND age >= ?"        "" ""
//	"updated_at > ?"               lastWeek
//	"created_at BETWEEN ? AND ?"   lastWeek today
//	"amount > (?)"                 db.Table("orders").Select("AVG(amount)")
func Where(condition string, args ...any) *gorm.DB {
	return db.Where(condition, args...)
}

// Not
//
// condition 和 args
//
//	"name = ?"                     ""
//	"name <> ?"                    ""
//	"name IN ?"                    []string{""}
//	"name LIKE ?"                  "%jin%"
//	"name = ? AND age >= ?"        "" ""
//	"updated_at > ?"               lastWeek
//	"created_at BETWEEN ? AND ?"   lastWeek today
func Not(condition string, args ...any) *gorm.DB {
	return db.Not(condition, args...)
}

// Or
//
// condition 和 args
//
//	"name = ?"                     ""
//	"name <> ?"                    ""
//	"name IN ?"                    []string{""}
//	"name LIKE ?"                  "%jin%"
//	"name = ? AND age >= ?"        "" ""
//	"updated_at > ?"               lastWeek
//	"created_at BETWEEN ? AND ?"   lastWeek today
func Or(condition string, args ...any) *gorm.DB {
	return db.Or(condition, args...)
}

// Order
//
// field
//
//	"name"
//	"age desc"
//	"name, age desc"
func Order(field string) *gorm.DB {
	return db.Order(field)
}

// Limit 查询多个数据时，限制返回最大数量
//
// limit
//
//	10
func Limit(limit int) *gorm.DB {
	return db.Limit(limit)
}

// Offset 结合排序和Limit可实现分页
//
// offset
//
//	10 跳过10条记录
//	-1 取消跳过跳过
func Offset(offset int) *gorm.DB {
	return db.Offset(offset)
}

// Select 选择输出记录的属性
//
// query
//
//	"name" "age"
//	[]string{"name", "age"}
//	"sum(age) as total"
//	"users.name, emails.email"
func Select(query any, args ...any) *gorm.DB {
	return db.Select(query, args...)
}

func Group() {
}

func Having() {
}

// Distinct 返回多条记录是，去掉重复的记录
//
// field
//
//	"name"
func Distinct(field ...any) *gorm.DB {
	return db.Distinct(field)
}

// Count 获取查询匹配的记录数
func Count(count *int64) {
	db.Count(count)
}

// Joins
//
//	左链接：保留左边所有数据，右边没匹配就为空
//	右链接：……
//	全链接：左右2边的行都保留
//
// query
//
//	"left join emails on emails.user_id = users.id"
//	"JOIN credit_cards ON credit_cards.user_id = users.id"
func Joins(query string, args ...interface{}) *gorm.DB {
	return db.Joins(query, args...)
}
