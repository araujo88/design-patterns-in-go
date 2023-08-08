# Data Access Object (DAO)

A Data Access Object (DAO) is a pattern used in software design to separate the logic that retrieves data from a database from the business logic of the application. Essentially, it creates an abstraction over the data source, which can be a database or an API or anything else where data is retrieved. The DAO interacts with the data source and provides an interface that the rest of the application can use, without having to know any details about how the data is stored or retrieved.

This pattern is widely used because it allows the underlying data source to be changed without affecting the rest of the application. This means that if you decide to change from one database system to another, you would only have to modify the DAO layer and not the entire application.

## Example

Here's a simple example of a DAO pattern in Go, working with a fictional user database.

First, let's define the User struct and the DAO interface:

```go
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
```

Now let's create a concrete implementation of this DAO interface. For simplicity, I'll use a mock in-memory database:

```go
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
```

Create a Service that Depends on the `UserDAO` Interface:

```go
type UserService struct {
    userDao UserDAO
}

func NewUserService(dao UserDAO) *UserService {
    return &UserService{userDao: dao}
}

func (s *UserService) GetUserDetails(id int) (*User, error) {
    return s.userDao.GetUser(id)
}
```

Use the Service in the Main Function:

```go
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
```

Here, the `UserService` depends on the `UserDAO` interface, not the concrete `UserDAOImpl`. This way, if you decide to change how users are stored (e.g., switching from an in-memory database to a SQL database), you can create a new implementation of the `UserDAO` interface without changing the `UserService`.

This example illustrates the DAO pattern's decoupling concept. By relying on interfaces, different parts of the application can interact without knowing the underlying implementation details, leading to more maintainable and flexible code.
