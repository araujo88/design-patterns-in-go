# Template

The "Template" design pattern defines the program skeleton of an algorithm in an algorithm class but delays some steps to subclasses. It lets one redefine certain steps of an algorithm without changing the algorithm's structure.

Basically, the pattern involves two main components:

 - Abstract Class: Defines abstract placeholder operations (primitive operations), concrete operations, and a template method.
 - Concrete Class: Implements the placeholder operations.

The template method in the abstract class defines the structure of an algorithm and calls placeholder operations as well as concrete operations.

## Example

The Template pattern is especially useful in software development scenarios where there's a clear sequence of steps to be performed, but the specifics of each step might vary.

Consider the example of a build tool for various project types. Suppose there's a general sequence of steps for building a software project:

 - Load configuration.
 - Fetch dependencies.
 - Compile source code.
 - Run tests.
 - Package for deployment.
 - While the sequence remains consistent, the specifics of each step can vary depending on the project type (e.g., a Java project vs. a Go project).

Let's demonstrate this with the Template pattern in Go:

```go
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
```

In this example:

 - BuildProcess serves as our "abstract class", defining the sequence for building software projects.
 - Build is our template method, defining the standard order in which these methods should be called.
 - JavaProject and GoProject are our concrete classes, implementing the steps of the build process for Java and Go projects, respectively.

This design makes it easy to introduce a build process for a new language or framework. Just implement the BuildProcess interface for the new project type and let the Build function handle the sequence.
