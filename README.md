# design-patterns-in-go

A list of design patterns implemented in Go with REAL WORLD examples!

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

### Behavioral Patterns

- **Observer**: Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
- **Command**: Encapsulates a request as an object, thereby letting you parameterize clients with queues, requests, and operations.
- **Strategy**: Defines a family of algorithms, encapsulates each one, and makes them interchangeable. Strategy lets the algorithm vary independently from the clients that use it.
- **State**: Allows an object to alter its behavior when its internal state changes. The object will appear to change its class.
- **Template**: Defines the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.
- **Iterator**: Provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
- **Mediator**: Defines an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and it lets you vary their interaction independently.
- **Memento**: Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later.

### Non-GoF Patterns

- **Service Locator**: Used to encapsulate the processes involved in obtaining a service with a strong abstraction layer.
- **Data Access Object (DAO)**: Used to separate the logic that retrieves data from a database from the business logic of the application.
