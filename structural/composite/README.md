# Composite

The Composite Design Pattern is a structural design pattern that allows you to compose objects into tree structures to represent part-whole hierarchies. The Composite pattern lets clients treat individual objects and compositions of objects uniformly.

This pattern is used when we need to treat a group of objects in the same way as a single instance of an object. The intent of a composite is to "compose" objects into tree structures to represent part-whole hierarchies. Implementing the composite pattern lets clients treat individual objects and compositions uniformly.

The composite pattern includes the following components:

- Component: It is an interface for objects in the composition. It has methods to manage the child components. It can be an abstract class which contains default implementations. If there is no specific default implementations to put in the abstract class, it can also be an interface.

- Leaf: It represents leaf objects in the composition. It has no children.

- Composite: It stores the child components and implements the child related operations in the component interface.

- Client: It manipulates objects in the composition through the component interface.

A real-world example could be an organization hierarchy. Every employee of an organization is an object. The organization consists of departments and each department also has a list of employees. Here, each department can be seen as a composite object as it can be further divided into sub-departments. This way, the organization is the composition of employees and departments. This can be efficiently represented using the composite design pattern.

## Example

Let's consider a file system as an example. In a file system, a directory can contain files or other directories. Here, both files and directories share similar operations such as calculating the size, searching for an item, etc.

In this example, the Component can be an interface named `FileSystemNode`, the Leaf can be a `File` and the Composite can be a `Directory`. The `Directory` can contain both `File` and `Directory`.

```go
// FileSystemNode : a component which can be a file or a directory
type FileSystemNode interface {
	GetSize() int
}

// File : a leaf node in the file system
type File struct {
	size int
}

// GetSize : returns size of the file
func (f *File) GetSize() int {
	return f.size
}

// Directory : a composite node in the file system
type Directory struct {
	children []FileSystemNode
}

// GetSize : returns total size of the directory (including all children)
func (d *Directory) GetSize() int {
	size := 0
	for _, child := range d.children {
		size += child.GetSize()
	}
	return size
}

// AddChild : adds a new child node to the directory
func (d *Directory) AddChild(child FileSystemNode) {
	d.children = append(d.children, child)
}
```

In this example, the client (`main` function) treats all files and directories uniformly by using the `FileSystemNode` interface. The client doesn't need to be aware of the concrete object it's dealing with. This is the primary benefit of the Composite Design Pattern.
