package database

import "gorm.io/gorm"

type ChatRoom struct {
	gorm.Model

	Users []User
}
