# Service locator

The Service Locator pattern is a design pattern used in software development to encapsulate the processes involved in obtaining a service with a strong abstraction layer. This pattern uses a central registry known as the "service locator", which on request returns the information necessary to perform a certain task.

Here is a brief summary of the pattern:

- The service locator is a known point for obtaining services.
- The service locator may 'know' how to get hold of the service itself, or it might contain a record of suitable services (i.e., a registry).
- An application will ask the service locator for the service by passing a string identifier or a key.
- The locator will return the service to the application so that the application can use it.

This pattern introduces an abstraction layer between the service consumer and the concrete classes providing those services. This abstraction layer provides a way to manage dependencies and to do runtime binding of the required service.

## Example

The Service Locator pattern is commonly used in situations where program configuration dictates what services to use. This pattern is particularly popular in microservice architectures and large enterprise systems, where services might be spread across different servers or even geographic regions.

The following example uses a service locator to switch between different databases (PostgreSQL and MySQL). We'll create a generic `DatabaseService` interface, and two structs implementing it, one for each database.

```go
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
```

In this example, we have defined a common interface (`DatabaseService`) for our databases and implemented it for PostgreSQL and MySQL. The `DatabaseServiceLocator` is our service locator which maintains a registry of services. We can fetch a specific database service by its name, and then call Connect() to simulate establishing a connection to the database.

We use the service locator in our main() function to fetch and connect to our databases. The type of database we're connecting to could be configured based on external factors, like application configuration or environment variables, making the system highly flexible.
