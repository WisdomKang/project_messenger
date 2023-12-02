package database

type UserService interface {
	CreateUser(user User)
	UpdateUser(user User)
	DeleteUser(user User)
}

type MessageService interface {
	CreateMessage(user User)
	UpdateMessage(user User)
	DeleteMessage(user User)
}

type ChatroomService interface {
	CreateChatroom(user User)
	UpdateChatroom(user User)
	DeleteChatroom(user User)
}
