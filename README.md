# design-patterns-in-go
A list of design patterns implemented in Go with REAL WORLD examples!

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
