package database

import "log"

type MyUserService struct {
}

func (service *MyUserService) CreateUser(user *User) {
	db := GetDB()
	result := db.Create(user)

	if result.Error != nil {
		log.Fatalf("Error on database service : %v", result.Error)
	}
}

func (service *MyUserService) GetUser(user *User) {
	db := GetDB()
	result := db.First(user)

	if result.Error != nil {
		log.Fatalf("Error on database service : %v", result.Error)
	}
}

func (service *MyUserService) UpdateUser(user *User) {
	db := GetDB()
	result := db.Save(user)

	if result.Error != nil {
		log.Fatalf("Error on database service : %v", result.Error)
	}

}

func (service *MyUserService) DeleteUser(user *User) {
	db := GetDB()

	result := db.Delete(user)

	if result.Error != nil {
		log.Fatalf("Error on database service : %v", result.Error)
	}
}

type MyMessageService struct {
}

func (service *MyMessageService) CreateMessage(user User) {}
func (service *MyMessageService) UpdateMessage(user User) {}
func (service *MyMessageService) DeleteMessage(user User) {}

type MyChatroomService struct {
}

func (service *MyChatroomService) CreateChatroom(user User) {}
func (service *MyChatroomService) UpdateChatroom(user User) {}
func (service *MyChatroomService) DeleteChatroom(user User) {}
