package main

import (
	"fmt"
	"sync"
)

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

func main() {
	service := NewCachedUserService(&RealUserService{})

	// Fetches from API
	service.GetUser("123")

	// Returns from cache
	service.GetUser("123")
}
