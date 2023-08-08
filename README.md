# design-patterns-in-go

A list of design patterns implemented in Go with REAL WORLD examples!

In a rush? Learn the top 10 design patterns:

## Top 10 design patterns according to ChatGPT

- **Singleton**: This pattern restricts a class from instantiating multiple objects. It is used where only a single instance of a class is required to control actions, such as logging, driver objects, caching, thread pool, database connections, and more.

- **Factory**: A creational pattern that provides one of the best ways to create an object. It provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

- **Abstract Factory**: This is a creational pattern that provides a way to encapsulate a group of individual factories that have a common theme without specifying their concrete classes.

- **Builder**: Used to construct a complex object step by step and the final step will return the object. This is useful when you need to create an object with many optional parameters or when the construction process needs to be explicit and controlled.

- **Prototype**: This pattern is used when creation of object directly is costly. For example, an object is to be created after a costly database operation. This pattern provides a mechanism to copy the original object.

- **Adapter**: This pattern works as a bridge between two incompatible interfaces. It combines capabilities of independent interfaces. A real-life example could be a case of a card reader, which acts as an adapter between memory card and a laptop.

- **Bridge:** This pattern is used to decouple an abstraction from its implementation so that the two can vary independently. This is useful when you want to avoid a permanent binding between an abstraction and its implementation.

- **Composite**: This pattern creates a tree structure of a group of objects. This pattern creates a class that contains a group of its own objects. This class provides ways to modify its group of the same objects.

- **Observer**: This pattern is used when there is a one-to-many relationship between objects such as if one object is modified, its dependent objects are to be notified automatically.

- **Strategy**: This pattern is used when we have multiple algorithm for a specific task and the client decides the actual implementation to be used at runtime.

## List of design patterns

### Creational Patterns

- **Singleton**: Ensures a class only has one instance and provides a global point of access to it.
- **Factory**: Provides an interface for creating objects in a super class, but allows subclasses to alter the type of objects that will be created.
- **Abstract Factory**: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.
- **Builder**: Separates the construction of a complex object from its representation so that the same construction process can create different representations.
- **Prototype**: Specifies the kind of objects to create using a prototypical instance and create new objects by copying this prototype.

### Structural Patterns

- **Adapter**: Converts the interface of a class into another interface the clients expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.
- **Decorator**: Attaches additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending functionality.
- **Facade**: Provides a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.
- **Composite**: Composes objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.
- **Proxy**: Provides a surrogate or placeholder for another object to control access to it.
- **Bridge**: Decouples an abstraction from its implementation so that the two can vary independently.
- **Flyweight**: Uses sharing to support large numbers of fine-grained objects efficiently.
- **Service Locator**: Used to encapsulate the processes involved in obtaining a service with a strong abstraction layer.

### Behavioral Patterns

- **Observer**: Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
- **Command**: Encapsulates a request as an object, thereby letting you parameterize clients with queues, requests, and operations.
- **Strategy**: Defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from the clients that use it.
- **State**: Allows an object to alter its behavior when its internal state changes. The object will appear to change its class.
- **Template**: Defines the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.
- **Iterator**: Provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
- **Mediator**: Defines an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and it lets you vary their interaction independently.
  Memento: Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later.
