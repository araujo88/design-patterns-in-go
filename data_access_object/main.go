package main

import (
	"fmt"
	"log"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

type UserDAO interface {
	GetUser(id int) (*User, error)
	GetAllUsers() ([]*User, error)
	SaveUser(user *User) error
	DeleteUser(id int) error
}

type UserDAOImpl struct {
	users map[int]*User
}

func (dao *UserDAOImpl) GetUser(id int) (*User, error) {
	user, exists := dao.users[id]
	if !exists {
		return nil, fmt.Errorf("User not found")
	}
	return user, nil
}

func (dao *UserDAOImpl) GetAllUsers() ([]*User, error) {
	userList := make([]*User, 0, len(dao.users))
	for _, user := range dao.users {
		userList = append(userList, user)
	}
	return userList, nil
}

func (dao *UserDAOImpl) SaveUser(user *User) error {
	dao.users[user.ID] = user
	return nil
}

func (dao *UserDAOImpl) DeleteUser(id int) error {
	_, exists := dao.users[id]
	if !exists {
		return fmt.Errorf("User not found")
	}
	delete(dao.users, id)
	return nil
}

type UserService struct {
	userDao UserDAO
}

func NewUserService(dao UserDAO) *UserService {
	return &UserService{userDao: dao}
}

func (s *UserService) GetUserDetails(id int) (*User, error) {
	return s.userDao.GetUser(id)
}

func main() {
	userDao := &UserDAOImpl{
		users: make(map[int]*User),
	}
	userService := NewUserService(userDao)

	user := &User{
		ID:        1,
		FirstName: "Giga",
		LastName:  "Chad",
	}
	userDao.SaveUser(user)

	retrievedUser, err := userService.GetUserDetails(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved User:", retrievedUser.FirstName, retrievedUser.LastName)
}
