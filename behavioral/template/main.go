package main

import "fmt"

// BuildProcess defines the steps of building a software project
type BuildProcess interface {
	LoadConfiguration()
	FetchDependencies()
	Compile()
	RunTests()
	Package()
}

// The template method
func Build(p BuildProcess) {
	p.LoadConfiguration()
	p.FetchDependencies()
	p.Compile()
	p.RunTests()
	p.Package()
}

type JavaProject struct{}

func (j *JavaProject) LoadConfiguration() {
	fmt.Println("Loading Maven configuration...")
}

func (j *JavaProject) FetchDependencies() {
	fmt.Println("Fetching Maven dependencies...")
}

func (j *JavaProject) Compile() {
	fmt.Println("Compiling Java source files...")
}

func (j *JavaProject) RunTests() {
	fmt.Println("Running JUnit tests...")
}

func (j *JavaProject) Package() {
	fmt.Println("Packaging into a JAR file...")
}

type GoProject struct{}

func (g *GoProject) LoadConfiguration() {
	fmt.Println("Loading Go modules configuration...")
}

func (g *GoProject) FetchDependencies() {
	fmt.Println("Fetching Go dependencies...")
}

func (g *GoProject) Compile() {
	fmt.Println("Building Go binary...")
}

func (g *GoProject) RunTests() {
	fmt.Println("Running Go tests...")
}

func (g *GoProject) Package() {
	fmt.Println("Packaging Go binary for deployment...")
}

func main() {
	java := &JavaProject{}
	Build(java)

	fmt.Println("---")

	golang := &GoProject{}
	Build(golang)
}
