package db

// CreateInBatches
// value可以是：
// []*user{}
// &[]User{}
func CreateInBatches(value any) {
	db.Create(value)
}
