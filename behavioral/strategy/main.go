package main

import (
	"fmt"
	"sort"
)

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

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	context := &SortingContext{}

	context.SetStrategy(&BubbleSortStrategy{})
	sortedData := context.ExecuteStrategy(data)
	fmt.Println("Bubble Sorted Data: ", sortedData)

	context.SetStrategy(&QuickSortStrategy{})
	sortedData = context.ExecuteStrategy(data)
	fmt.Println("Quick Sorted Data: ", sortedData)
}
