# Mediator

The Mediator design pattern is a behavioral design pattern that aims to reduce coupling between classes that communicate with each other. Instead of classes communicating directly with each other, they interact through a mediator object. This makes it easier to modify, add, or remove components without having to alter the other interacting components. The Mediator pattern is particularly useful in large and complex systems where you want to control the interactions between classes.

## Participants in the Mediator Pattern

- **Mediator**: The interface that declares methods to communicate with Colleagues.
- **ConcreteMediator**: Implements the Mediator interface and keeps track of all Colleagues.
- **Colleague**: An interface for components that will communicate with each other via the Mediator.
- **ConcreteColleague**: The class that will implement the Colleague interface and communicate via the Mediator.

## Example

Let's consider a chat room as an example, where multiple users can send messages to the chat room, and the chat room will display the messages to all users.

Here is a simple implementation in Go:

```go
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
```

In this example:

- `Mediator` is represented by the `Mediator` interface, and `ChatRoom` is the `ConcreteMediator`.
- `Colleague` is represented by the `Colleague` interface, and `User` is the `ConcreteColleague`.

The `ChatRoom` keeps track of all `User` objects and handles the forwarding of messages between them. Each `User` communicates only with the `ChatRoom`, reducing the coupling between `User` objects.
