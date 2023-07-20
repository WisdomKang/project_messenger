package database

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string
	Provider string
	Phone    string

	ChatRooms []ChatRoom
	Messages  []Message
}
