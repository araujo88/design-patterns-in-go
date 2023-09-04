# Visitor

The Visitor design pattern is a behavioral design pattern that allows you to add further operations to objects without modifying the classes of the objects you're working with. It allows you to define a new operation without altering the classes of the elements on which it operates. This pattern can be particularly useful when dealing with a structure composed of a mixture of different types of objects and you want to perform operations that depend on the concrete classes of these objects.

### Participants

- **Visitor**: This is an interface that declares a visit operation for each type of concrete element in the object structure.
- **ConcreteVisitor**: This class implements the Visitor interface and defines the operations to be performed on each type of element.
- **Element**: This is an interface that declares an `Accept` method that accepts a visitor.
- **ConcreteElement**: This class implements the Element interface and allows the visitor to visit it.

## Example

The Visitor pattern is often used in cases where you have a complex data structure, like an Abstract Syntax Tree (AST) in a compiler, or a Document Object Model (DOM) in a web browser. The Visitor pattern allows you to extend functionalities without modifying these structures.

An Abstract Syntax Tree (AST) is a tree representation of the syntactic structure of source code written in a programming language. Each node in the tree denotes a construct in the language. It's "abstract" in the sense that it omits certain syntax rules that aren't necessary for understanding the program's structure. For example, grouping parentheses might be omitted because the structure of the tree makes the execution order of operations clear.

### Components

- **Node**: Each node represents a language construct. For example, a node could represent an arithmetic operation like addition, a control flow construct like a `for` loop, or a data construct like an array.
  
- **Edge**: Edges in an AST represent the relationship between the nodes. For example, an edge from an "If statement" node to a "Boolean expression" node signifies that the Boolean expression is a condition for the If statement.

- **Root**: The root of the tree represents the entire program or the start symbol in the grammar of the programming language. 

### Use Cases

1. **Compilers**: The front end of a compiler usually transforms source code into an AST as an intermediate step in translating code into machine language or bytecode.
  
2. **Code Analysis**: Tools that perform static code analysis often use ASTs to understand the structure of the code.

3. **Refactoring**: Code refactoring tools use ASTs to accurately transform the structure of the code.

4. **Syntax Highlighting**: Some advanced text editors use ASTs for better syntax highlighting.


Consider a simple arithmetic expression `1 + 2 * 3`. An AST representing this expression might look like this:

```
        [+]
       /   \
      [1]  [*]
          /   \
         [2]  [3]
```

Here, the root node represents the `+` operation. The left child is `1` and the right child is another operation `*`. This tree structure makes it clear that `2 * 3` should be evaluated first, and then added to `1`, adhering to the rules of arithmetic.

Let's consider a simplified version of an AST. Assume that we have two types of expressions: `Literal` expressions (just numbers) and `Addition` expressions (which add two other expressions together). Now, imagine we want to perform multiple operations on these expressions: evaluate them to get a result and convert them to a string in infix notation.

Here's how you could use the Visitor pattern to implement this in Go:

```go
// Expression is the 'Element' interface
type Expression interface {
	Accept(Visitor) string
}

// Visitor interface
type Visitor interface {
	VisitLiteral(*Literal) string
	VisitAddition(*Addition) string
}

// Literal is a 'ConcreteElement'
type Literal struct {
	value int
}

func (l *Literal) Accept(visitor Visitor) string {
	return visitor.VisitLiteral(l)
}

// Addition is another 'ConcreteElement'
type Addition struct {
	left, right Expression
}

func (a *Addition) Accept(visitor Visitor) string {
	return visitor.VisitAddition(a)
}

// Evaluator is a 'ConcreteVisitor' that evaluates expressions
type Evaluator struct{}

func (e *Evaluator) VisitLiteral(l *Literal) string {
	return strconv.Itoa(l.value)
}

func (e *Evaluator) VisitAddition(a *Addition) string {
	left := a.left.Accept(e)
	right := a.right.Accept(e)
	return fmt.Sprintf("(%s + %s)", left, right)
}

// Renderer is another 'ConcreteVisitor' that converts expressions to strings
type Renderer struct{}

func (r *Renderer) VisitLiteral(l *Literal) string {
	return strconv.Itoa(l.value)
}

func (r *Renderer) VisitAddition(a *Addition) string {
	left := a.left.Accept(r)
	right := a.right.Accept(r)
	return fmt.Sprintf("(%s + %s)", left, right)
}
```

In this example:

- `Expression` is the Element interface that declares an `Accept` method to accept a visitor.
- `Literal` and `Addition` are Concrete Elements that represent literal numbers and addition operations, respectively.
- `Visitor` is the Visitor interface that declares `VisitLiteral` and `VisitAddition` methods.
- `Evaluator` and `Renderer` are Concrete Visitors. `Evaluator` evaluates expressions, and `Renderer` converts them to strings.

This example is a simplified AST, but the Visitor pattern makes it easy to add new operations (like optimization, code generation, etc.) without changing the `Expression`, `Literal`, or `Addition` classes. It's particularly useful in software engineering scenarios like compilers and data transformation pipelines.
