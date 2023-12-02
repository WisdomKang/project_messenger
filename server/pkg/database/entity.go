package database

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string
	Provider string
	Phone    string

	ChatRooms []ChatRoom `gorm:"many2many:user_chatroom"`
	Messages  []Message
}

type ChatRoom struct {
	gorm.Model
	Users    []User `gorm:"many2many:user_chatroom"`
	Messages []Message
}

type Message struct {
	gorm.Model

	Contents string
	Type     string

	// 채팅방과 연결된 외래키
	ChatRoomID uint

	// 사용자와 연결된 외래키
	UserID uint
}
