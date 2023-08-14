# Command

The command design pattern is a behavioral design pattern used to encapsulate a request as an object, thereby allowing for parameterization of clients with queues, requests, and operations. It also allows for the decoupling of the sender and receiver objects. The Command Pattern consists of the following main components:

 - Command: This defines an interface for executing an operation.
 - ConcreteCommand: This is a subclass of Command and implements the execute method by invoking the corresponding operation on the receiver object.
 - Client: The Client creates a ConcreteCommand object and sets its receiver.
 - Invoker: It asks the command to carry out the request.
 - Receiver: It knows how to perform the operation to satisfy a request.

The command pattern can be particularly useful when you need to allow for undo functionality, queuing of commands, task scheduling or logging of command operations. It also provides greater flexibility in terms of organizing code and adding new commands without modifying existing code.

## Example

Here's a practical example that could be part of a text editor, where we implement a simple undo functionality using the command pattern in Go.

```go
// Command interface
type Command interface {
	Execute()
	Unexecute()
}

// Receiver
type TextEditor struct {
	text string
}

func (t *TextEditor) Add(s string) {
	t.text += s
}

func (t *TextEditor) Delete(n int) {
	if n <= len(t.text) {
		t.text = t.text[:len(t.text)-n]
	}
}

func (t *TextEditor) Display() {
	fmt.Println(t.text)
}

// ConcreteCommand for Adding Text
type AddTextCommand struct {
	editor *TextEditor
	text   string
}

func (c *AddTextCommand) Execute() {
	c.editor.Add(c.text)
}

func (c *AddTextCommand) Unexecute() {
	c.editor.Delete(len(c.text))
}

// Invoker
type CommandInvoker struct {
	history []Command
}

func (i *CommandInvoker) Execute(c Command) {
	c.Execute()
	i.history = append(i.history, c)
}

func (i *CommandInvoker) Undo() {
	if len(i.history) > 0 {
		lastCommand := i.history[len(i.history)-1]
		lastCommand.Unexecute()
		i.history = i.history[:len(i.history)-1]
	}
}
```

In this example, the text editor is our receiver, and we have a specific command for adding text (`AddTextCommand`) with both Execute and `Unexecute` methods. The invoker (`CommandInvoker`) manages the execution of commands and keeps a history of executed commands, allowing for simple undo functionality.

By employing the command pattern, we've made it easy to extend the text editor with additional commands (like copy, paste, formatting) without changing existing code. The undo functionality is also neatly encapsulated within the invoker, allowing for more advanced features like multi-level undo or redo with minimal changes.
