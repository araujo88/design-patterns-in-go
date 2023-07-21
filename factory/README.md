# Factory

The Factory Method Design Pattern is a creational pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. This pattern is used when a class cannot anticipate the class of objects it needs to create.

The Factory pattern is often used in situations where a system should be independent of how its objects are created, composed, and represented. It provides a simple way of decoupling the concrete objects from the parts of the system that use them.

A factory method in a class is a method that returns objects of a varying prototype or class from some method call, which is assumed to be "new". The factories are a great way to implement the concept of abstraction by encapsulating the responsibilities of creating and initializing complex objects.

## Example

In this example, we have a `Logger` interface that different types of loggers (like `ConsoleLogger` and `FileLogger`) implement. The `LoggerFactory` can create different types of loggers depending on the argument passed to `CreateLogger`.

This could be part of a larger logging library where you can choose between different types of loggers depending on your requirements. For example, you might use a `ConsoleLogger` in a development environment to print logs for debugging, but in a production environment, you might use a `FileLogger` or maybe a `DatabaseLogger` to store logs more permanently.
