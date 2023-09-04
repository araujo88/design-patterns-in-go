package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func main() {
	context := &Context{}
	expression := context.Parse(strings.Fields("3 4 + 2 * 7 /"))

	fmt.Println("Result:", expression.Interpret())
}
