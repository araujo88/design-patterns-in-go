# Dependency Injection

Dependency Injection is a design pattern that allows the removal of hard-coded dependencies and makes it possible to change them, whether at run-time or compile-time. This pattern is particularly useful for decoupling components and improving code maintainability and testability.

In simple terms, Dependency Injection means that if a class needs access to an external resource (like a database, or another class), instead of creating it inside the class, you "inject" it into the class from the outside.

### Participants

- **Client**: The class that depends on the service.
- **Service**: The service that the client needs.
- **Injector**: This class injects the service into the client.

## Example

Below is a simplified example in Go that demonstrates Dependency Injection. We have an `Author` struct that uses an interface `Writer` to write something. We then have two implementations of `Writer`: `ConsoleWriter` and `FileWriter`. Through dependency injection, we can easily switch between writing to the console and writing to a file without modifying the `Author` struct.

```go
// Writer interface, the 'Service' in our example
type Writer interface {
	Write([]byte) error
}

// ConsoleWriter writes to the console, one of the 'ConcreteService'
type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(data []byte) error {
	_, err := fmt.Println(string(data))
	return err
}

// FileWriter writes to a file, another 'ConcreteService'
type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(data []byte) error {
	return ioutil.WriteFile(fw.filename, data, 0644)
}

// Author is the 'Client'
type Author struct {
	writer Writer // Author depends on a Writer to do the writing
}

// NewAuthor creates a new author with a writer
func NewAuthor(writer Writer) *Author {
	return &Author{writer}
}

// WriteArticle uses the injected writer to write an article
func (a *Author) WriteArticle(article string) error {
	return a.writer.Write([]byte(article))
}
```

In this example:

- `Writer` is the service interface that defines a `Write` method.
- `ConsoleWriter` and `FileWriter` are concrete implementations of the `Writer` interface.
- `Author` is the client that depends on a service implementing the `Writer` interface.
  
We inject the `Writer` dependency into `Author` via the `NewAuthor` function. This decouples `Author` from the concrete implementations (`ConsoleWriter` and `FileWriter`) and makes the system easier to extend, test, and maintain.

With Dependency Injection, you can easily swap out the `Writer` implementation without modifying the `Author` code, making the system more modular and easier to test.
