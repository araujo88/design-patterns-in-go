package main

import "fmt"

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

func main() {
	editor := &TextEditor{}
	invoker := &CommandInvoker{}

	addCommand1 := &AddTextCommand{editor: editor, text: "Hello, "}
	invoker.Execute(addCommand1)
	editor.Display() // Output: Hello,

	addCommand2 := &AddTextCommand{editor: editor, text: "world!"}
	invoker.Execute(addCommand2)
	editor.Display() // Output: Hello, world!

	invoker.Undo()
	editor.Display() // Output: Hello,
}
