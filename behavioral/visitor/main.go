package main

import (
	"fmt"
	"strconv"
)

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

func main() {
	// Create an AST for the expression: (1 + (2 + 3))
	expression := &Addition{
		left:  &Literal{value: 1},
		right: &Addition{left: &Literal{value: 2}, right: &Literal{value: 3}},
	}

	// Evaluate the expression
	evaluator := &Evaluator{}
	result := expression.Accept(evaluator)
	fmt.Printf("Evaluation: %s\n", result)

	// Render the expression
	renderer := &Renderer{}
	result = expression.Accept(renderer)
	fmt.Printf("Rendering: %s\n", result)
}
