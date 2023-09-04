package main

import "fmt"

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
