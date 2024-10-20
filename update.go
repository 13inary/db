package main

// Save
//
//	不能和Model()一起使用
//	字段是默认值也会修改数据库
//	若数据不包含主键：则执行create
//	若数据包含主键：执行update
//
// table
//
//	&user{}
func Save(table any) error {
	result := db.Save(table)
	return result.Error
}

// Update
//
//	更新单列
//	结合Model()使用，若没有指定主键，则更新所有匹配的记录
func Update(column string, value interface{}) error {
	result := db.Update(column, value)
	return result.Error
}

// UpdateColumn
//
//	类似Update
//	跳过hook方法、不追踪更新时间
func UpdateColumn() {
}

// Updates
//
//	更新多列
//	结合Model()使用，若没有指定主键，则更新所有匹配的记录
//	忽略默认值字段
//	使用select可以选择更新的字段，若"*"则表示更新所有字段
//	使用omit可以选择不更新的字段
func Updates(table any) error {
	result := db.Updates(table)
	return result.Error
}

// UpdateColumns
//
//	类似Updates
//	跳过hook方法、不追踪更新时间
func UpdateColumns() {
}
