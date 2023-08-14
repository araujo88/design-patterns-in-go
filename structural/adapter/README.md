# Adapter

The Adapter Design Pattern is a structural design pattern that allows two incompatible interfaces to work together. This incompatibility could be because of differences in method names, argument types, or other differences in the interfaces.

The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.

Here's an example: Suppose you have a legacy system that returns data in an older format (like XML), and you have a new system that expects the data to be in a different format (like JSON). Instead of rewriting the legacy system, you could use an Adapter. This Adapter could take the XML data, convert it to JSON, and then provide the data to the new system.

In a real-world scenario, an adapter is like having a USB-C to USB-A converter. You have a device that only supports USB-C as an input but you have a USB-A connector cable. So you use a converter/adapter to solve the incompatibility.

The Adapter Design Pattern is all about creating an intermediary abstraction that translates, or adapts, the old component to the new system. An adapter wraps the legacy system with a new interface so that the legacy system can be used by the new system through the new interface. This way, the new system only interacts with the Adapter but the legacy system doesn't need to be modified.

## Example

Imagine we have a legacy system that logs messages in a certain way and a new system that has a different interface for logging.

The legacy system has a `LegacyLogger` with a `LogMessage` method:

```go
type LegacyLogger struct {}

func (l *LegacyLogger) LogMessage(message string) {
    fmt.Println("Legacy logger: " + message)
}
```

The new system uses an interface called `NewLogger` with a `Log` method:

```go
type NewLogger interface {
    Log(message string)
}

type NewSystemLogger struct {}

func (l *NewSystemLogger) Log(message string) {
    fmt.Println("New system logger: " + message)
}
```

You want the new system to be able to use the `LegacyLogger` without changing its code. You can create an adapter that implements the `NewLogger` interface, but uses the `LegacyLogger` internally:

```go
type LoggerAdapter struct {
    legacyLogger *LegacyLogger
}

func (a *LoggerAdapter) Log(message string) {
    // Adapting the interface here:
    a.legacyLogger.LogMessage(message)
}
```

Now, you can use `LegacyLogger` in the new system:

```go
func logMessage(message string, logger NewLogger) {
    logger.Log(message)
}

func main() {
    legacyLogger := &LegacyLogger{}
    loggerAdapter := &LoggerAdapter{legacyLogger: legacyLogger}

    logMessage("Hello, world!", loggerAdapter)
}
```

In this example, `LoggerAdapter` is the adapter. It wraps `LegacyLogger` (the adaptee) and makes it compatible with `NewLogger` (the target interface). The `Log` method in `LoggerAdapter` adapts the interface of `LegacyLogger` to the `NewLogger` interface. The client code (the `logMessage` function) works with the `NewLogger` interface, so it can now use `LegacyLogger` through the `LoggerAdapter`. This is the essence of the Adapter Pattern.
