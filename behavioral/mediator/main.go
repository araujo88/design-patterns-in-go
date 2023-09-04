package main

import "fmt"

// Mediator interface
type Mediator interface {
	SendMessage(string, Colleague)
}

// Colleague interface
type Colleague interface {
	Send(string)
	Receive(string)
}

// ConcreteMediator
type ChatRoom struct {
	users []Colleague
}

func (c *ChatRoom) Register(user Colleague) {
	c.users = append(c.users, user)
}

func (c *ChatRoom) SendMessage(message string, sender Colleague) {
	for _, user := range c.users {
		if user != sender {
			user.Receive(message)
		}
	}
}

// ConcreteColleague
type User struct {
	name    string
	chatMed Mediator
}

func NewUser(name string, chatMed Mediator) *User {
	return &User{name: name, chatMed: chatMed}
}

func (u *User) Send(message string) {
	fmt.Printf("%s sends message: %s\n", u.name, message)
	u.chatMed.SendMessage(message, u)
}

func (u *User) Receive(message string) {
	fmt.Printf("%s received message: %s\n", u.name, message)
}

func main() {
	chatroom := &ChatRoom{}

	alice := NewUser("Alice", chatroom)
	bob := NewUser("Bob", chatroom)
	charlie := NewUser("Charlie", chatroom)

	chatroom.Register(alice)
	chatroom.Register(bob)
	chatroom.Register(charlie)

	alice.Send("Hi, everyone!")
	bob.Send("Hello, Alice!")
	charlie.Send("Hey, how are you?")
}
