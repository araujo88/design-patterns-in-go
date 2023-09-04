# Chain of Responsibility

The Chain of Responsibility design pattern is a behavioral pattern that allows you to pass a request along a chain of handlers. Each handler decides either to process the request or to pass it along the chain. This pattern allows you to decouple the sender and receiver components, making the code more modular and easier to maintain.

In the Chain of Responsibility pattern, usually, each handler has a reference to the next handler in the chain. A handler decides whether it can handle the request; if it can, it does so, otherwise, it delegates the request to the next handler in the chain.

### Participants in the Chain of Responsibility Pattern

- **Handler**: The interface that declares the method for handling requests.
- **ConcreteHandler**: The class that implements the Handler interface and performs the actual work, or delegates the work to the next handler in the chain.
- **Client**: The class that initiates the request to a handler on the chain.

## Example

Let's consider a simple example where we have different levels of logging: `Info`, `Warning`, and `Error`. Each logging level is handled by a different logger. However, if one logger can't handle a particular log level, it passes the request to the next logger in the chain.

Here is how you can implement this in Go:

```go
// Log levels
const (
	Info = iota
	Warning
	Error
)

// Handler interface
type Logger interface {
	SetNext(Logger)
	LogMessage(int, string)
}

// ConcreteHandler
type ConsoleLogger struct {
	level int
	next  Logger
}

func (c *ConsoleLogger) SetNext(next Logger) {
	c.next = next
}

func (c *ConsoleLogger) LogMessage(level int, message string) {
	if c.level <= level {
		fmt.Printf("Writing to console: %s\n", message)
	}
	if c.next != nil {
		c.next.LogMessage(level, message)
	}
}

// Another ConcreteHandler
type ErrorLogger struct {
	level int
	next  Logger
}

func (e *ErrorLogger) SetNext(next Logger) {
	e.next = next
}

func (e *ErrorLogger) LogMessage(level int, message string) {
	if e.level <= level {
		fmt.Printf("Writing to error log: %s\n", message)
	}
	if e.next != nil {
		e.next.LogMessage(level, message)
	}
}

func main() {
	// Setting up the chain
	consoleLogger := &ConsoleLogger{level: Info}
	errorLogger := &ErrorLogger{level: Error}

	consoleLogger.SetNext(errorLogger)

	// Making requests
	consoleLogger.LogMessage(Info, "This is an informational message.")
	consoleLogger.LogMessage(Error, "This is an error message.")
}
```

In this example:

- `Logger` is the `Handler` interface that defines the method for logging messages.
- `ConsoleLogger` and `ErrorLogger` are `ConcreteHandlers`. Each decides whether it can handle the given logging level or pass it to the next handler in the chain.
- The `main()` function acts as the `Client` that initiates the requests.

Run the code, and you'll see that depending on the log level, the appropriate logger(s) handle the message.

With the Chain of Responsibility pattern, you can easily add more loggers in the chain without modifying the existing code, thus following the Open/Closed Principle.
