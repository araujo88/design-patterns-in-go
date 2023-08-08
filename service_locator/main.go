package main

import (
	"fmt"
	"log"
)

// Define our Service Interface
type DatabaseService interface {
	GetName() string
	Connect() string
}

type PostgreSQL struct{}

func (p *PostgreSQL) GetName() string {
	return "PostgreSQL"
}

func (p *PostgreSQL) Connect() string {
	return "Connected to PostgreSQL"
}

type MySQL struct{}

func (m *MySQL) GetName() string {
	return "MySQL"
}

func (m *MySQL) Connect() string {
	return "Connected to MySQL"
}

// Service Locator
type DatabaseServiceLocator struct {
	services map[string]DatabaseService
}

func (sl *DatabaseServiceLocator) GetService(serviceName string) (DatabaseService, error) {
	service, ok := sl.services[serviceName]
	if ok {
		return service, nil
	}
	return nil, fmt.Errorf("Service Not Found")
}

func (sl *DatabaseServiceLocator) RegisterService(service DatabaseService) {
	sl.services[service.GetName()] = service
}

func main() {
	locator := &DatabaseServiceLocator{
		services: make(map[string]DatabaseService),
	}

	postgres := &PostgreSQL{}
	locator.RegisterService(postgres)

	mysql := &MySQL{}
	locator.RegisterService(mysql)

	// Fetch PostgreSQL service and connect
	service, err := locator.GetService("PostgreSQL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(service.Connect())

	// Fetch MySQL service and connect
	service, err = locator.GetService("MySQL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(service.Connect())
}
