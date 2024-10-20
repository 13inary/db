package db

// 若表包含gorm.DeletedAt字段，那么下面删除都是软删除。

// Delete
//
//	若指定主键：删除单条记录
//	若没有指定主键：批量删除
//
// tables
//
//	&user{}或&[]user{}
func Delete(tables any) error {
	result := db.Delete(tables)
	return result.Error
}

// Unscoped
//
//	若在查询前使用：可以查询到软删除的记录
//	若在删除前使用：可以彻底删除记录
func Unscoped() {
}
