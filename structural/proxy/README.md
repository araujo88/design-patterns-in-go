# Proxy

The proxy design pattern provides a surrogate or placeholder for another object to control access to it. It involves a class, called the proxy, which represents the functionality of another class. The proxy pattern is useful for tasks like lazy instantiation (creating an object only when it's actually needed), permission control, and more sophisticated memory management.

Here's a simple breakdown of how the proxy pattern works:

 - The Subject defines a common interface for the RealObject and the Proxy so that the Proxy can be used anywhere the RealObject is expected.
 - The RealObject defines the real object that the proxy represents.
 - The Proxy maintains a reference that lets the proxy access the real object. It handles requests and forwards these requests to the real object.

 ## Example

A practical example could be a system that interacts with an external service API. Without a proxy, every part of your code that needed to communicate with the API would have to know how to do so. With a proxy, you encapsulate the communication logic in one place, making the code cleaner and easier to maintain.

Let's say you have a service that retrieves user information from a third-party API. You may want to implement a caching mechanism to avoid unnecessary network requests if the same data is requested multiple times. You can implement this with the proxy pattern.

```go
type User struct {
	ID   string
	Name string
}

type UserService interface {
	GetUser(ID string) (*User, error)
}

type RealUserService struct{}

func (s *RealUserService) GetUser(ID string) (*User, error) {
	// Simulate a real API request
	fmt.Println("Fetching user from API...")
	return &User{ID: ID, Name: "John Doe"}, nil
}

type CachedUserService struct {
	realService UserService
	cache       map[string]*User
	mux         sync.Mutex
}

func (s *CachedUserService) GetUser(ID string) (*User, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if user, ok := s.cache[ID]; ok {
		fmt.Println("Returning user from cache...")
		return user, nil
	}
	user, err := s.realService.GetUser(ID)
	if err != nil {
		return nil, err
	}
	s.cache[ID] = user
	return user, nil
}

func NewCachedUserService(service UserService) UserService {
	return &CachedUserService{
		realService: service,
		cache:       make(map[string]*User),
	}
}
```

In this example, `UserService` is the interface for the user service, `RealUserService` fetches user data from the API, and `CachedUserService` is the proxy that adds caching. If the requested user data is in the cache, `CachedUserService` returns it without calling the real service. Otherwise, it fetches the data from `RealUserService`, stores it in the cache, and then returns it.
