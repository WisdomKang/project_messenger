package database

import "gorm.io/gorm"

type Message struct {
	gorm.Model

	Contents string
	Type     string

	Chatroom ChatRoom

	User User
}
