# Builder

The Builder Design Pattern is a creational pattern that is used to construct a complex object step by step. It separates the construction of an object from its representation so that the same construction process can create different representations.

When you are building a complex object that has multiple parts or configuration options, creating and configuring the object directly within client code or via a constructor can become unwieldy and hard to read. The Builder pattern addresses this issue by providing a way to build the complex object step by step, with each step clearly encapsulated by a method in the builder.

## Example

The Builder pattern is most useful when dealing with a complex object that requires a significant or complex process to construct. It's often employed in contexts where objects are built from a variety of parts that need to be created and assembled. It's also particularly helpful when you might have multiple representations of an object, or when you want to encapsulate the creation logic separate from the main application.

In this example, `RequestBuilder` is used to construct a complex `http.Request` object. Each Set method in the builder sets a different attribute of the request and returns the builder itself, allowing the methods to be chained. The `Build` method is finally used to construct the `http.Request` object. This approach can make the creation of complex request objects more manageable and readable.
