# Strategy

The Strategy design pattern is a behavioral design pattern that lets the algorithm vary independently from the clients that use it. It defines a family of algorithms, encapsulates each one, and makes them interchangeable.

The strategy pattern involves using an interface to abstractly define an operation that needs to be carried out. Concrete implementations of the interface will encapsulate different strategies or ways to carry out the operation.

The object that uses the strategy pattern will maintain a reference to a Strategy object and will communicate with this object only through the strategy interface.

The advantage of the strategy pattern is that it promotes algorithmic independence. It decouples the behavior of an object from the object itself.

Here's the general structure of Strategy design pattern:

- Strategy (interface): This is an interface used to abstractly define the operation(s) that we're planning to implement with different strategies.

- ConcreteStrategy (class): This is a class that implements the Strategy interface, providing its own version of the interface method(s).

- Context (class): This is a class that uses a strategy to do something. It interacts with a strategy through the Strategy interface only, meaning it's decoupled from the ConcreteStrategy.

This pattern can be beneficial when we need to select an algorithm at runtime. It's also helpful when we have a lot of similar classes that only differ in their behaviors. Instead of implementing many versions of your class with different behaviors, you can separate the behaviors as strategies and use the same context class in all versions.

## Example

Let's use the Strategy pattern to select between a Bubble sort and a Quick sort strategy at runtime:

```go
// SortStrategy defines the method for sorting
type SortStrategy interface {
	Sort([]int) []int
}

// BubbleSortStrategy uses bubble sort algorithm for sorting
type BubbleSortStrategy struct{}

func (b *BubbleSortStrategy) Sort(input []int) []int {
	n := len(input)
	sorted := make([]int, n)
	copy(sorted, input)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	return sorted
}

// QuickSortStrategy uses quick sort algorithm for sorting
type QuickSortStrategy struct{}

func (q *QuickSortStrategy) Sort(input []int) []int {
	sorted := make([]int, len(input))
	copy(sorted, input)
	sort.Ints(sorted)
	return sorted
}

// SortingContext is the context that uses the strategy
type SortingContext struct {
	strategy SortStrategy
}

func (s *SortingContext) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *SortingContext) ExecuteStrategy(input []int) []int {
	return s.strategy.Sort(input)
}
```

In the example above, `SortStrategy` is the strategy interface, `BubbleSortStrategy` and `QuickSortStrategy` are concrete strategies, and `SortingContext` is the context. Depending on what we set as the `SortStrategy`, the appropriate sorting algorithm will be applied.
