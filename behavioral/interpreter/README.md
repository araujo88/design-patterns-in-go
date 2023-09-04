# Interpreter

The Interpreter design pattern is a behavioral design pattern that provides a way to evaluate language grammar for particular languages. It involves defining a representation for the language's grammar and an interpreter to interpret the grammar. The pattern is particularly useful for building compilers and interpreters, but it's generally less commonly used than other design patterns.

### Participants

- **Abstract Expression**: Declares an abstract Interpret operation that is common to all nodes in the abstract syntax tree.
- **Terminal Expression**: Implements the Interpret operation for terminal expressions.
- **Nonterminal Expression**: Implements the Interpret operation for grammar rules which include other expressions.
- **Context**: Contains the global information that is interpreted.

## Example

Let's consider a simple arithmetic problem where we want to evaluate an expression given in Reverse Polish Notation (RPN). In RPN, every operator follows all of its operands. For example, the expression `3 4 +` evaluates to `7`.

Here is how you can implement it in Go:

```go
// Expression interface
type Expression interface {
	Interpret() int
}

// TerminalExpression
type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

// NonTerminalExpression
type Operation struct {
	left, right Expression
	operator    string
}

func (o *Operation) Interpret() int {
	switch o.operator {
	case "+":
		return o.left.Interpret() + o.right.Interpret()
	case "-":
		return o.left.Interpret() - o.right.Interpret()
	case "*":
		return o.left.Interpret() * o.right.Interpret()
	case "/":
		return o.left.Interpret() / o.right.Interpret()
	}
	return 0
}

// Context
type Context struct {
	stack []Expression
}

func (c *Context) Parse(tokens []string) Expression {
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			right := c.stack[len(c.stack)-1]
			c.stack = c.stack[:len(c.stack)-1]
			left := c.stack[len(c.stack)-1]
			c.stack = c.stack[:len(c.stack)-1]

			c.stack = append(c.stack, &Operation{
				left:     left,
				right:    right,
				operator: token,
			})
		} else {
			num, _ := strconv.Atoi(token)
			c.stack = append(c.stack, &Number{value: num})
		}
	}
	return c.stack[len(c.stack)-1]
}
```

In this example:

- `Expression` is the Abstract Expression that declares an `Interpret()` method.
- `Number` is a Terminal Expression that represents individual numbers.
- `Operation` is a Non-terminal Expression that represents arithmetic operations.
- `Context` stores the global information (in this case, the operand stack), and it also has a `Parse` method to build the abstract syntax tree for the given RPN expression.


The Interpreter pattern lets you build a parse tree from your language of choice and then evaluate it. However, it's worth noting that this pattern can lead to a large number of classes and can be hard to manage for complex grammars. In such cases, alternative techniques such as parser combinators or parser generators may be more appropriate.
