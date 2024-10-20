package main

import "gorm.io/gorm"

func Session() *gorm.DB {
	return db.Session(&gorm.Session{})
}
