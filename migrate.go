package main

// Migrate
//
// tables
//
//	&user{}
func Migrate(tables ...any) error {
	return db.AutoMigrate(tables...)
}
