# Iterator

The "Iterator" design pattern provides a way to access the elements of an aggregate object (like a collection) sequentially without exposing its underlying representation. The main purpose of this pattern is to decouple the iteration logic from the collection itself.

The Iterator pattern typically involves at least three components:

 - Iterator: This defines an interface for accessing and traversing elements.
 - ConcreteIterator: This implements the Iterator interface, keeping track of the current position in the traversal of the aggregate.
 - Aggregate: This defines an interface for creating an Iterator object.
 - ConcreteAggregate: This implements the Aggregate interface and returns an instance of the ConcreteIterator.

## Example

Iterators are often used in real-world software systems to abstract the process of iterating over different types of collections. For a more practical example, let's consider a scenario where we have a software system managing different types of data sources. One common task might be to iterate over records from each data source, regardless of how that data is stored or represented.

**Scenario**: Let's say you're building a system that manages user profiles. The profiles can be stored in different data structures: a slice (array-like structure in Go) and a map. You want to have a consistent way to iterate over user profiles, irrespective of the underlying data structure.

**Implementation**:

1. **UserProfile**: Represents the user profile.
2. **Aggregate**: An interface to create an Iterator.
3. **Iterator**: Interface to traverse through elements.

**Concrete Aggregates**:
1. **UserProfileSlice**: Represents a slice of user profiles.
2. **UserProfileMap**: Represents a map of user profiles.

**Concrete Iterators**:
1. **SliceIterator**: Iterates over a slice.
2. **MapIterator**: Iterates over a map.

Here's how you can implement the above:

```go
type UserProfile struct {
	ID   int
	Name string
}

// Aggregate interface
type UserProfileCollection interface {
	CreateIterator() Iterator
}

// Iterator interface
type Iterator interface {
	HasNext() bool
	Next() *UserProfile
}

// Concrete Aggregate: Slice
type UserProfileSlice struct {
	profiles []UserProfile
}

func (us *UserProfileSlice) CreateIterator() Iterator {
	return &SliceIterator{profiles: us.profiles, index: 0}
}

type SliceIterator struct {
	profiles []UserProfile
	index    int
}

func (si *SliceIterator) HasNext() bool {
	return si.index < len(si.profiles)
}

func (si *SliceIterator) Next() *UserProfile {
	if si.HasNext() {
		profile := &si.profiles[si.index]
		si.index++
		return profile
	}
	return nil
}

// Concrete Aggregate: Map
type UserProfileMap struct {
	profiles map[int]UserProfile
	keys     []int
}

func (um *UserProfileMap) CreateIterator() Iterator {
	keys := make([]int, 0, len(um.profiles))
	for k := range um.profiles {
		keys = append(keys, k)
	}
	return &MapIterator{profiles: um.profiles, keys: keys, index: 0}
}

type MapIterator struct {
	profiles map[int]UserProfile
	keys     []int
	index    int
}

func (mi *MapIterator) HasNext() bool {
	return mi.index < len(mi.keys)
}

func (mi *MapIterator) Next() *UserProfile {
	if mi.HasNext() {
		key := mi.keys[mi.index]
		profile := mi.profiles[key]
		mi.index++
		return &profile
	}
	return nil
}
```

The main advantage here is that the code iterating over user profiles remains the same, irrespective of whether the underlying data structure is a slice or a map. This decouples the iteration logic from the data structure, making the code more modular and easier to maintain.
