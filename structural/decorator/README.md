# Decorator

The Decorator design pattern allows us to add new functionalities to existing objects without altering their structures. This pattern involves a set of decorator classes that are used to wrap concrete components. Decorator classes mirror the type of the components they decorate but add or override behavior.

The key feature of this pattern is the ability to "wrap" or "decorate" an object multiple times, with different decorators, adding cumulative behavior to the original object. The pattern is particularly useful for adding responsibilities to objects dynamically and optionally.

Key Points:

 - Component Interface: This defines the interface for objects that can have responsibilities added to them dynamically.
 - Concrete Component: It is the object to which new responsibilities are added.
 - Decorator: This maintains a reference to a Component object and defines an interface that conforms to the Component's interface.
 - Concrete Decorators: These extend the Decorator and add specific responsibilities.

## Example

Let's consider an example with input/output (I/O) operations in a system. Imagine you want to read data from a file. At its core, you just read raw bytes. But often you'll want to enhance this process with additional features such as:

 - Buffering: Read large chunks at once to improve performance.
 - Decompression: If your file is compressed, you might want to decompress it on-the-fly while reading.
 - Encryption: Maybe your data is encrypted, and you need to decrypt it when reading.

Each of these features can be seen as a "decoration" to the basic file reading operation. Using the decorator pattern, you can dynamically construct a chain of operations for file reading.

Here's an example in Go, focusing on the structure of the decorator pattern:

```go
// Reader is our base component interface
type Reader interface {
	Read() string
}

// FileReader reads from a file
type FileReader struct{}

func (fr *FileReader) Read() string {
	return "data from file"
}

// Buffered reader as a decorator
type BufferedReader struct {
	reader Reader
}

func (br *BufferedReader) Read() string {
	data := br.reader.Read()
	return "buffered " + data
}

// Decrypt reader as another decorator
type DecryptReader struct {
	reader Reader
}

func (dr *DecryptReader) Read() string {
	data := dr.reader.Read()
	return "decrypted " + data
}
```

In this example:

 - `FileReader` provides basic file reading functionality.
 - `BufferedReader` and `DecryptReader` are decorators that extend the base reading functionality.

By using the decorator pattern, we can stack or chain these operations in any combination we want, providing great flexibility.
